package main

import (
	"fmt"
)

// Implements: Expandable, Form

type Macro struct {
	Args []Symbol
	Body *List
}

func (m Macro) Eval(e Environment) (Form, error) {
	return m, nil
}

func (m Macro) Expand(args []Form) (Form, error) {
	argmap := make(map[string]Form)

	for idx, a := range args {
		argmap[m.Args[idx].Name] = a
	}

	return m.Body.Expand(argmap)
}

func (m Macro) Call(e Environment, args []Form) (Form, error) {
	if len(args) > len(m.Args) {
		return nil, fmt.Errorf("Too many arguments")
	}

	x, err := m.Expand(args)
	if err != nil {
		return nil, err
	}

	return x.Eval(e)
}
