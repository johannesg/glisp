package main

import (
	"fmt"
	"log"
)

type LispFunction func(Environment, []Form) (Form, error)

var builtIns = map[string]BuiltInFunction{
	"+":   BuiltInFunction{Fn: BuiltInAdd},
	"add": BuiltInFunction{Fn: BuiltInAdd},
	// "=":    BuiltInFunction{Fn: BuiltInEquality,
	"eval": BuiltInFunction{Fn: BuiltInEval},
	"def":  BuiltInFunction{Fn: BuiltInDef},
	"fn":   BuiltInFunction{Fn: BuiltInFn},
	"vars": BuiltInFunction{Fn: BuiltInVars},
}

func (f BuiltInFunction) Invoke(e Environment, args []Form) (Form, error) {
	return f.Fn(e, args)
}

func (f UserFunction) Invoke(e Environment, args []Form) (Form, error) {
	if len(args) > len(f.Args) {
		return nil, fmt.Errorf("Too many arguments")
	}

	local := NewEnvironment(e)

	for idx, a := range args {
		local.SetVar(f.Args[idx].Name, a)
	}

	return f.Body.Eval(local)
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

func BuiltInFn(e Environment, args []Form) (Form, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("fn: Wrong number of arguments")
	}

	var fargs *Vector
	var fbody *List
	var ok bool

	if fargs, ok = args[0].(*Vector); !ok {
		return nil, fmt.Errorf("fn: first argument must be a vector")
	}

	if fbody, ok = args[1].(*List); !ok {
		return nil, fmt.Errorf("fn: second argument must be a list")
	}

	symbols := make([]Symbol, len(fargs.Items))
	for idx, fa := range fargs.Items {
		if s, ok := fa.(Symbol); ok {
			symbols[idx] = s
		} else {
			return nil, fmt.Errorf("fn: all arguments must evaluate to symbols")
		}
	}

	f := UserFunction{
		Args: symbols,
		Body: fbody,
	}
	return f, nil
}

func BuiltInVars(e Environment, args []Form) (Form, error) {
	// items := []Form{}

	for k, v := range e.Vars() {
		log.Printf("%v: %#v", k, v)
	}

	return nil, nil
}

// func BuiltInEquality(args []Form) (Form, error) {
// 	if len(args) == 0 {
// 		return nil, fmt.Errorf("Wrong number of arguments")
// 	}

// 	for a := range args {

// 	}
// }
