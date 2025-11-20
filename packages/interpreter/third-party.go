package interpreter

import (
	"log"
	"reflect"

	"github.com/traefik/yaegi/interp"
	"github.com/traefik/yaegi/stdlib"
)

type yeagiInterp struct {
	interp *interp.Interpreter
}

func newYeagiInterp() *yeagiInterp {
	r := &yeagiInterp{
		interp: interp.New(interp.Options{}),
	}

	r.interp.Use(stdlib.Symbols)

	for _, pkg := range defaultImports {
		if _, err := r.Eval(`import `+pkg); err != nil {
			log.Fatalf("Failed to import default packages. Failed to load %s: %v", pkg, err)
		}
	}

	return r
}

func (i *yeagiInterp) Eval(src string) (reflect.Value, error) {
	return i.interp.Eval(src)
}
