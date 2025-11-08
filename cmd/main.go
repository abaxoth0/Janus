package main

import (
  "bufio"
  "fmt"
  "os"
  "github.com/traefik/yaegi/interp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	i := interp.New(interp.Options{})

	fmt.Println("Janus - simple REPL for Golang. Type \"/exit\" to exit")

	fmt.Print(">>> ")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			fmt.Print(">>> ")
			continue
		}

		if line == "/exit" {
			break
		}

		val, err := i.Eval(line)
		if err != nil {
			fmt.Println("Error:", err)
		} else if val.IsValid() {
			fmt.Println(val)
		}

		fmt.Print(">>> ")
	}
}
