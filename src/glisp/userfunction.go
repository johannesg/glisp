package main

import (
	"fmt"
)

type UserFunction struct {
	Args []Symbol
	Body *List
}

func (f UserFunction) Eval(e Environment) (Form, error) {
	return f, nil
}

func (f UserFunction) Invoke(e Environment, args []Form) (Form, error) {
	if len(args) > len(f.Args) {
		return nil, fmt.Errorf("Too many arguments")
	}

	local := NewEnvironment(e)

	for idx, a := range args {
		if da, err := a.Eval(e); err == nil {
			local.SetVar(f.Args[idx].Name, da)
		} else {
			return nil, err
		}
	}

	return f.Body.Eval(local)
}
