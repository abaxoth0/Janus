package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

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

		if isCmd(line) {
			res := execCmd(line)
			if res.Err != nil {
				if res.Err == errInvalidCmd {
					fmt.Printf("Invalid command: \"%s\" (type \"/help\" to see available commands)\n", line)
					continue
				}
				return res.Err
			}
			if res.ShouldBreak {
				break
			}
			if res.ShouldContinue {
				continue
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
