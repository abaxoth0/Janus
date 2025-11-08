package repl

import (
	"bufio"
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

func (r *REPL) Run(interp interpreter.Interpreter) {
	fmt.Print(">>> ")
	for r.scanner.Scan() {
		line := r.scanner.Text()
		if line == "" {
			fmt.Print(">>> ")
			continue
		}

		if line == "/exit" {
			break
		}

		val, err := interp.Eval(line)
		if err != nil {
			fmt.Println("Error:", err)
		} else if val.IsValid() {
			fmt.Println(val)
		}

		fmt.Print(">>> ")
	}
}
