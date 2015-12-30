package main

import (
	"fmt"
	"log"
)

type LispFunction func(Environment, []Form) (Form, error)

var builtIns = map[string]BuiltInFunction{
	"+":   BuiltInFunction{Fn: BuiltInAdd, EvalArgs: true},
	"add": BuiltInFunction{Fn: BuiltInAdd, EvalArgs: true},
	// "=":    BuiltInFunction{Fn: BuiltInEquality,
	"do":       BuiltInFunction{Fn: BuiltInDo, EvalArgs: true},
	"def":      BuiltInFunction{Fn: BuiltInDef, EvalArgs: true},
	"defn":     BuiltInFunction{Fn: BuiltInDefn},
	"defmacro": BuiltInFunction{Fn: BuiltInDefmacro},
	"fn":       BuiltInFunction{Fn: BuiltInFn},
	"vars":     BuiltInFunction{Fn: BuiltInVars},
}

func (f BuiltInFunction) Invoke(e Environment, args []Form) (Form, error) {
	if f.EvalArgs {
		eargs := make([]Form, len(args))
		for idx, a := range args {
			if ea, err := a.Eval(e); err == nil {
				eargs[idx] = ea
			} else {
				return nil, err
			}
		}
		return f.Fn(e, eargs)
	}
	return f.Fn(e, args)
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

func BuiltInDo(e Environment, args []Form) (ret Form, err error) {
	for _, a := range args {
		ret, err = a.Eval(e)
		if ret != nil {
			return
		}
	}
	return
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

func BuiltInDefn(e Environment, args []Form) (Form, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("defn: Wrong number of arguments")
	}

	s, ok := args[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("defn: First argument must be a symbol")
	}

	var fargs *Vector
	var fbody *List

	if fargs, ok = args[1].(*Vector); !ok {
		return nil, fmt.Errorf("defn: first argument must be a vector")
	}

	if fbody, ok = args[2].(*List); !ok {
		return nil, fmt.Errorf("defn: second argument must be a list")
	}

	if fn, err := BuiltInFn(e, []Form{fargs, fbody}); err != nil {
		return nil, err
	} else {
		e.SetVar(s.Name, fn)
		return fn, nil
	}
}

func BuiltInDefmacro(e Environment, args []Form) (Form, error) {
	return nil, nil
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
