package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/abaxoth0/Janus/packages/ascii"
	"github.com/abaxoth0/Janus/packages/interpreter"
)

type REPL struct {
	scanner *bufio.Scanner
	cursor 	*cursor
	inputReader *bufio.Reader
	inputBuf	[]rune
}

func New(r io.Reader) *REPL {
	return &REPL{
		scanner: bufio.NewScanner(r),
		cursor: new(cursor),
		inputReader: bufio.NewReader(os.Stdin),
	}
}

func (r *REPL) readln() (string, error) {
	r.inputBuf = r.inputBuf[:0]

	for {
		char, _, err := r.inputReader.ReadRune()
		if err != nil {
			return "", err
		}

		if char == ascii.LineFeed || char == ascii.CarriageReturn {
			r.cursor.FlushLine()
			break
		}
		if ascii.IsControlChar(char) {
			continue
		}

		if char == ascii.Backspace {
			if len(r.inputBuf) > 0 {
				r.inputBuf = r.inputBuf[:len(r.inputBuf)-1]
			}
			continue
		}

		r.inputBuf = append(r.inputBuf, char)
		r.cursor.WriteChar(char)
	}

	return string(r.inputBuf), nil
}

func (r *REPL) Run(interp interpreter.Interpreter) error {
	for {
		fmt.Print("\n>>> ")

		line, err := r.readln()
		if err != nil {
			return err
		}
		if line == "" {
			continue
		}

		if r.isCmd(line) {
			res := r.exec(line)
			if res.Err != nil {
				if res.Err == errInvalidCmd {
					fmt.Println("Invalid command: " + line)
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
			fmt.Print("Error:", err)
		}
		if val.IsValid() {
			fmt.Print(val)
		}
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
