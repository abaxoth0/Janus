package repl

import (
	"errors"
	"fmt"
	"strings"
)

func isCmd(s string) bool {
	return s[0] == '/'
}

const (
	exitCmd string = "/exit"
	typeCmd string = "/type"
	helpCmd string = "/help"
)

type cmdResult struct {
	Code		   string
	Err            error
	ShouldBreak    bool
	ShouldContinue bool
}

var errInvalidCmd = errors.New("invalid command")

func execCmd(cmd string) *cmdResult {
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
				Err: errors.New("Invalid command format, use case example: /type <value/identifier>"),
			}
		}
		return &cmdResult{
			Code: `fmt.Sprintf("%T",`+cmdValue+`);`,
		}
	case helpCmd:
		fmt.Print(cmdHelpMessage)
		return &cmdResult{
			ShouldContinue: true,
		}
	}
	return &cmdResult{
		Err: errInvalidCmd,
	}
}

const cmdHelpMessage string =
`Available commands:
 	/help - show this message
	/exit - exit this programm
	/type <value/identifier> - show type of the given value/identifier
`
