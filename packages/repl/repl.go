package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"

	"github.com/abaxoth0/Janus/packages/interpreter"
)

type REPL struct {
	scanner *bufio.Scanner
}

func New(r io.Reader) *REPL {
	return &REPL{
		scanner: bufio.NewScanner(r),
	}
}

func (r *REPL) Run(interp interpreter.Interpreter) error {

	fmt.Print(">>> ")
	for r.scanner.Scan() {
		line := r.scanner.Text()
		if line == "" {
			fmt.Print(">>> ")
			continue
		}

		if r.isCmd(line) {
			res := r.exec(line)
			if res.Err != nil {
				if res.Err == errInvalidCmd {
					fmt.Println("Invalid command: " + line)
					fmt.Print(">>> ")
					continue
				}
				return res.Err
			}
			if res.ShouldBreak {
				break
			}
		}

		val, err := interp.Eval(line)
		if err != nil {
			fmt.Println("Error:", err)
		} else if val.IsValid() {
			fmt.Println(val)
		}

		fmt.Print(">>> ")
	}

	return nil
}

func (r *REPL) isCmd(s string) bool {
	return s[0] == '/'
}

const (
	exitCmd string = "/exit"
)

type cmdResult struct {
	Err 		error
	ShouldBreak bool
}

var errInvalidCmd = errors.New("invalid command")

func (r *REPL) exec(cmd string) *cmdResult {
	switch cmd {
	case exitCmd:
		return &cmdResult{
			ShouldBreak: true,
		}
	}
	return &cmdResult{
		Err: errInvalidCmd,
	}
}
