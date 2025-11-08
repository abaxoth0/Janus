package interpreter

import (
	"reflect"

	"github.com/traefik/yaegi/interp"
)

type yeagiInterp struct {
	interp *interp.Interpreter
}

func newYeagiInterp() *yeagiInterp {
	return &yeagiInterp{
		interp: interp.New(interp.Options{}),
	}
}

func (i *yeagiInterp) Eval(src string) (reflect.Value, error) {
	return i.interp.Eval(src)
}
