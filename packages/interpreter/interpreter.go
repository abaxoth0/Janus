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

var defaultImports []string = []string{
	`"fmt"`,
	`"os"`,
	`"io"`,
	`"strings"`,
	`"strconv"`,
	`"bytes"`,
	`"net/http"`,
	`"encoding/json"`,
	`"context"`,
	`"sync"`,
	`"testing"`,
	`"time"`,
	`"regexp"`,
}

var ErrInvalidInterpreterType = errors.New("invalid interpreter type")

func New(t Type) (Interpreter, error) {
	switch t {
	case ThirdParty:
		return newYeagiInterp(), nil
	}
	return nil, ErrInvalidInterpreterType
}
