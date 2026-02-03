package repl

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/abaxoth0/Janus/packages/interpreter"
	"github.com/chzyer/readline"
)

const InputPrompt string = ">>> "

type REPL struct {
	rl    *readline.Instance
}

func New(r io.Reader) *REPL {
	rl, err := readline.NewEx(&readline.Config{
		Prompt: InputPrompt,
		HistoryFile: "/tmp/janus_history",
		HistoryLimit: 1000,
		InterruptPrompt: "^C",
	})
	if err != nil {
		log.Fatalln(err)
	}

	return &REPL{
		rl: rl,
	}
}

func (r *REPL) read() (string, error) {
	indentLevel := 0
	lines := []string{}

	for {
		if len(lines) > 0 {
			indent := strings.Repeat(" ", indentLevel*4)
			r.rl.SetPrompt(indent + "... ")
		} else {
			r.rl.SetPrompt(InputPrompt)
		}

		line, err := r.rl.Readline()
		if err != nil {
			return "", err
		}

		indentLevel += strings.Count(line, "{")
		indentLevel -= strings.Count(line, "}")

		lines = append(lines, line)

		if r.shouldContinue(lines) {
			continue
		}

		break
	}

	return strings.Join(lines, "\n"), nil
}

var operators []string = []string{"+", "-", "*", "/", "=", "==", "!=", ">", "<", "&&", "||"}

func (r *REPL) shouldContinue(lines []string) bool {
	if len(lines) == 0 {
		return false
	}

	code := strings.Join(lines, "\n")
	trimmedLastLine := strings.TrimSpace(lines[len(lines)-1])

	if strings.HasSuffix(strings.TrimSpace(trimmedLastLine), "\\") {
		return true
	}

	if strings.Count(code, "{") > strings.Count(code, "}") {
		return true
	}

	if strings.Count(code, "(") > strings.Count(code, ")") {
		return true
	}

	for _, op := range operators {
		if strings.HasSuffix(trimmedLastLine, op) {
			return true
		}
	}

	return false
}

func (r *REPL) Run(interp interpreter.Interpreter) error {
	rl, err := readline.New(InputPrompt)
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		r.rl.SetPrompt(InputPrompt)

		input, err := r.read()
		if err != nil {
			if err == io.EOF {
				break
			}
			if err == readline.ErrInterrupt {
				continue
			}
			return err
		}
		if input == "" {
			continue
		}

		if isCmd(input) {
			res := runCmd(input)
			if res.Err != nil {
				if res.Err == errUnknownCmd {
					fmt.Printf("Unknown command: \"%s\" (type \"/help\" to see available commands)\n", input)
					continue
				}
				fmt.Println("Error:", res.Err.Error())
				continue
			}
			if res.ShouldBreak {
				break
			}
			if res.ShouldContinue {
				continue
			}
			if res.Code != "" {
				input = res.Code
			}
		}

		val, err := interp.Eval(input)
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
		if val.IsValid() {
			fmt.Println(val)
		}
	}

	return nil
}
