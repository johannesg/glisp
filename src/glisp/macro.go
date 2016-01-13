package main

import (
	"fmt"
)

// Implements: Expandable, Form

type Macro struct {
	Args *ArgVector
	Body *List
}

func (m Macro) Eval(e Environment) (Form, error) {
	return m, nil
}

func (m Macro) Expand(args []Form) (Form, error) {
	local := NewEnvironment(nil)
	if !m.Args.Match(args, local) {
		return nil, fmt.Errorf("Could not match arguments")
	}

	return m.Body.Expand(local)
}

func (m Macro) Call(e Environment, args []Form) (Form, error) {
	x, err := m.Expand(args)
	if err != nil {
		return nil, err
	}

	return x.Eval(e)
}
