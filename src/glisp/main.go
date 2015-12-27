package main

import (
	"fmt"

	"gopkg.in/readline.v1"
)

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	env := NewEnvironment(nil)

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		ret, err := env.Eval(line)

		if err != nil {
			println(fmt.Sprintf("%v", err))
		} else {
			println(fmt.Sprintf("%v", ret))
		}
	}
}
