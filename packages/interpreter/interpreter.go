package interpreter

import (
	"errors"
	"reflect"
)

type Interpreter interface {
	Eval(src string) (reflect.Value, error)
}

type Type uint16

const (
	ThirdParty Type = 1 << iota
)

var ErrInvalidInterpreterType = errors.New("invalid interpreter type")

func New(t Type) (Interpreter, error) {
	switch t {
	case ThirdParty:
		return newYeagiInterp(), nil
	}
	return nil, ErrInvalidInterpreterType
}
