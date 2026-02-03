package repl

import (
	"errors"
	"fmt"
	"strings"
)

func isCmd(s string) bool {
	return s[0] == '/'
}

type cmdType uint8

const (
	exitApp cmdType = iota+1
	showType
	pkgImport
	showHelp
)

func (t cmdType) String() string {
	switch t {
	case exitApp:
		return "/exit"
	case showType:
		return "/type"
	case pkgImport:
		return "/import"
	case showHelp:
		return "/help"
	}
	panic("unknown cmd type")
}

type cmdKind uint8

const (
	plain cmdKind = iota+1
	targeted
)

type cmd interface {
	Type() cmdType
	Kind() cmdKind
}

type plainCmd = cmdType

func (c plainCmd) Type() cmdType {
	return c
}

func (c plainCmd) Kind() cmdKind {
	return plain
}

type targetedCmd struct {
	target string

	plainCmd
}

func (c targetedCmd) Kind() cmdKind {
	return targeted
}

type cmdResult struct {
	Code		   string
	Err            error
	ShouldBreak    bool
	ShouldContinue bool
}

var errUnknownCmd = errors.New("uknown command")
var errMissingCmdTarget = errors.New("missing target for command")
var errTooManyCmdTargets = errors.New("too many targets for command (redundant whitespace?)")

func parseCmd(rawCmd string) (cmd, error) {
	splitCmd := strings.Split(rawCmd, " ")

	rawCmdType := splitCmd[0]

	var _type cmdType
	kind := plain

	switch rawCmdType {
	case exitApp.String():
		_type = exitApp
	case showHelp.String():
		_type = showHelp
	case showType.String():
		_type = showType
		kind = targeted
	case pkgImport.String():
		_type = pkgImport
		kind = targeted
	default:
		return nil, errUnknownCmd
	}

	if kind == plain {
		return plainCmd(_type), nil
	}

	if len(splitCmd) == 1 {
		return nil, errMissingCmdTarget
	}
	if len(splitCmd) > 2 {
		return nil, errTooManyCmdTargets
	}

	return targetedCmd{
		target: splitCmd[1],
		plainCmd: plainCmd(_type),
	}, nil
}

func convertToTargetedCmd(cmd cmd) targetedCmd {
	r, ok := cmd.(targetedCmd)
	if !ok {
		panic(fmt.Sprintf("Invalid data type, expected targetedCmd, but got %T", cmd))
	}
	return r
}

func runCmd(rawCmd string) *cmdResult {
	cmd, err := parseCmd(rawCmd)
	if err != nil {
		return &cmdResult{Err: err}
	}

	switch cmd.Type() {
	case exitApp:
		return &cmdResult{
			ShouldBreak: true,
		}
	case showType:
		cmd := convertToTargetedCmd(cmd)
		return &cmdResult{
			Code: `fmt.Sprintf("%T",`+cmd.target+`);`,
		}
	case pkgImport:
		cmd := convertToTargetedCmd(cmd)
		packages := strings.Split(cmd.target, ",")
		if len(packages) == 1 {
			return &cmdResult{ Code: `import "`+cmd.target+`"` }
		}
		return &cmdResult{
			Code: "import (\n\""+strings.Join(packages, "\"\n\"")+"\"\n)",
		}

	case showHelp:
		fmt.Print(cmdHelpMessage)
		return &cmdResult{
			ShouldContinue: true,
		}
	}
	panic("UNREACHABLE")
}

const cmdHelpMessage string =
`Available commands:
 	/help - show this message
	/exit - exit this programm
	/type <value/identifier> - show type of the given value/identifier
	/import <packages> - import comma-separated package (you also may use Go imports directly in code)
`
