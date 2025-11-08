package main

import (
	"fmt"
	"os"

	"github.com/abaxoth0/Janus/packages/interpreter"
	"github.com/abaxoth0/Janus/packages/repl"
)

func main() {
	fmt.Println("Janus — simple REPL (Read-Eval-Print Loop) for Golang. Type \"/exit\" to exit")

	interp, err := interpreter.New(interpreter.ThirdParty)
	if err != nil {
		panic(err)
	}

	REPL := repl.New(os.Stdin)

	REPL.Run(interp)
}
