package main

import (
	"fmt"
	"glisp/reader"

	"gopkg.in/readline.v1"
)

func main() {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()
	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		r := reader.NewReader(line)
		for {
			f, err := r.Read()
			if f == nil {
				break
			}

			if err != nil {
				println(fmt.Sprintf("%v", err))
			}

			r, err := f.Eval()

			if err != nil {
				println(fmt.Sprintf("%v", err))
			} else {
				println(fmt.Sprintf("%v", r))
			}

		}
	}
}
