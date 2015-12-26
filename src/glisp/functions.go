package main

import (
	"fmt"
)

type LispFunction func(Environment, []Form) (Form, error)

var builtIns = map[string]LispFunction{
	"+":   BuiltInAdd,
	"add": BuiltInAdd,
	// "=":    BuiltInEquality,
	"eval": BuiltInEval,
	"def":  BuiltInDef,
}

func (e *environment) Invoke(s Symbol, args []Form) (Form, error) {
	f, ok := builtIns[s.Name]
	if ok {
		return f(e, args)
	}

	v, ok := e.vars[s.Name]
	if !ok {
		return nil, fmt.Errorf("Var '%v' not found", s.Name)
	}

	return v.Eval(e)
}

func BuiltInAdd(e Environment, args []Form) (Form, error) {
	result := 0
	for _, a := range args {

		n, ok := a.(Number)
		if !ok {
			return nil, fmt.Errorf("Expected number")
		}

		result += n.Val
	}
	return Number{Val: result}, nil
}

func BuiltInEval(e Environment, args []Form) (Form, error) {
	q, ok := args[0].(*QForm)
	if !ok {
		return args[0].Eval(e)
	}

	return q.Form.Eval(e)
}

func BuiltInDef(e Environment, args []Form) (Form, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("def: Wrong number of arguments")
	}

	s, ok := args[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("def: First argument must be a symbol")
	}

	if len(args) < 2 {
		return s, nil
	}

	r, err := args[1].Eval(e)
	if err != nil {
		return nil, err
	}

	e.SetVar(s.Name, r)
	return s, nil
}

// func BuiltInEquality(args []Form) (Form, error) {
// 	if len(args) == 0 {
// 		return nil, fmt.Errorf("Wrong number of arguments")
// 	}

// 	for a := range args {

// 	}
// }
