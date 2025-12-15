package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/abaxoth0/Janus/packages/interpreter"
	"github.com/chzyer/readline"
)

const InputPrefx string = ">>> "

type REPL struct {
	scanner     *bufio.Scanner
	inputReader *bufio.Reader
	inputBuf    []rune
}

func New(r io.Reader) *REPL {
	return &REPL{
		scanner:     bufio.NewScanner(r),
		inputReader: bufio.NewReader(os.Stdin),
	}
}

func (r *REPL) Run(interp interpreter.Interpreter) error {
	rl, err := readline.New(InputPrefx)
	if err != nil {
		return err
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		if line == "" {
			continue
		}

		if r.isCmd(line) {
			res := r.exec(line)
			if res.Err != nil {
				if res.Err == errInvalidCmd {
					fmt.Printf("Invalid command: \"%s\"\n", line)
					continue
				}
				return res.Err
			}
			if res.ShouldBreak {
				break
			}
			if res.Code != "" {
				line = res.Code
			}
		}

		val, err := interp.Eval(line)
		if err != nil {
			fmt.Println("Error:" + err.Error())
		}
		if val.IsValid() {
			fmt.Println(val)
		}
	}

	return nil
}

func (r *REPL) isCmd(s string) bool {
	return s[0] == '/'
}

const (
	exitCmd string = "/exit"
	typeCmd string = "/type"
)

type cmdResult struct {
	Err         error
	ShouldBreak bool
	Code		string
}

var errInvalidCmd = errors.New("invalid command")

func (r *REPL) exec(cmd string) *cmdResult {
	splitCmd := strings.Split(cmd, " ")

	var cmdType  string = splitCmd[0]
	var cmdValue string

	switch cmdType {
	case exitCmd:
		return &cmdResult{
			ShouldBreak: true,
		}
	case typeCmd:
		cmdValue = splitCmd[1]
		if len(splitCmd) != 2 {
			return &cmdResult{
				Err: errors.New("Invalid command format, use case example: /type <identifier>"),
			}
		}
		return &cmdResult{
			Code: `fmt.Sprintf("%T",`+cmdValue+`);`,
		}
	}
	return &cmdResult{
		Err: errInvalidCmd,
	}
}
