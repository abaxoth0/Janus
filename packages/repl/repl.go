package repl

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"

	ansix364 "github.com/abaxoth0/Janus/packages/ansix3.64"
	"github.com/abaxoth0/Janus/packages/ascii"
	"github.com/abaxoth0/Janus/packages/interpreter"
	"golang.org/x/term"
)

const InputPrefx string = "\r\n>>> "

type REPL struct {
	scanner     *bufio.Scanner
	cursor      *cursor
	inputReader *bufio.Reader
	inputBuf    []rune
}

func New(r io.Reader) *REPL {
	return &REPL{
		scanner:     bufio.NewScanner(r),
		cursor:      nil,
		inputReader: bufio.NewReader(os.Stdin),
	}
}

func (r *REPL) handleEscSeq(seq string) {
	switch seq {
	case ansix364.Left:
		r.cursor.Back()
	case ansix364.Right:
		if r.cursor.GetX() < len(r.inputBuf)-1+len(InputPrefx) {
			r.cursor.Forward()
		}
	default:
		// panic(fmt.Sprintf("unsupported escape sequence: %q", seq))
	}
}

func (r *REPL) readln() (string, error) {
	r.inputBuf = r.inputBuf[:0]

	const (
		stateNormal = iota
		stateEscape
		stateCSI
	)

	state := stateNormal
	csiBuffer := ""

	for {
		char, _, err := r.inputReader.ReadRune()
		if err != nil {
			return "", err
		}

		switch state {
		case stateNormal:
			switch {
			case char == ascii.Escape:
				state = stateEscape
			case char == ascii.LineFeed || char == ascii.CarriageReturn:
				return string(r.inputBuf), nil
			case char == ascii.Backspace:
				if len(r.inputBuf) > 0 {
					r.inputBuf = r.inputBuf[:len(r.inputBuf)-1]
					r.cursor.FlushChar()
				}
			case ascii.IsControlChar(char):
				// ignore other (unhandled) control chars
			default:
				r.inputBuf = append(r.inputBuf, char)
				r.cursor.WriteChar(char)
			}
		case stateEscape:
			if char == ansix364.CSIPrefix {
				state = stateCSI
				csiBuffer = ""
			} else {
				// Reset if not a CSI sequence
				// (currently there are no other supported escape sequences)
				state = stateNormal
			}
		case stateCSI:
			csiBuffer += string(char)

			// CSI sequences end with a letter command
			if ascii.IsAlpha(char) {
				fullSeq := "\033[" + csiBuffer
				r.handleEscSeq(fullSeq)
				csiBuffer = ""
				state = stateNormal
			} else if len(csiBuffer) > 10 { // Safety limit
				state = stateNormal
				csiBuffer = ""
			}
		}
	}
}

func (r *REPL) Run(interp interpreter.Interpreter) error {
	stdin := int(os.Stdin.Fd())

	oldState, err := term.MakeRaw(stdin)
	if err != nil {
		return err
	}
	defer term.Restore(stdin, oldState)

	r.cursor = newCursor(len(InputPrefx) - 1).Rewind()

	for {
		fmt.Print(InputPrefx)

		line, err := r.readln()
		if err != nil {
			return err
		}
		if line == "" {
			r.cursor.Rewind()
			continue
		}

		if r.isCmd(line) {
			res := r.exec(line)
			if res.Err != nil {
				if res.Err == errInvalidCmd {
					r.cursor.Writeln("Invalid command: " + line)
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
			r.cursor.Writeln("Error:" + err.Error())
		}
		if val.IsValid() {
			r.cursor.NewLine().Rewind().RawWrite(val)
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
	Err         error
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
