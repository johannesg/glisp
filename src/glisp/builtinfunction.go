package main

import (
	"fmt"
	"log"
)

type LispFunction func(Environment, []Form) (Form, error)

type BuiltInFunction struct {
	Fn       LispFunction
	EvalArgs bool
}

func (f BuiltInFunction) Eval(e Environment) (Form, error) {
	return f, nil
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

var builtIns = map[string]BuiltInFunction{
	"+":   BuiltInFunction{Fn: BuiltInAdd, EvalArgs: true},
	"add": BuiltInFunction{Fn: BuiltInAdd, EvalArgs: true},
	// "=":    BuiltInFunction{Fn: BuiltInEquality,
	"do":          BuiltInFunction{Fn: BuiltInDo, EvalArgs: true},
	"def":         BuiltInFunction{Fn: BuiltInDef, EvalArgs: true},
	"defn":        BuiltInFunction{Fn: BuiltInDefn},
	"defmacro":    BuiltInFunction{Fn: BuiltInDefmacro},
	"fn":          BuiltInFunction{Fn: BuiltInFn},
	"vars":        BuiltInFunction{Fn: BuiltInVars},
	"macroexpand": BuiltInFunction{Fn: BuiltInMacroExpand},
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
	if len(args) < 3 {
		return nil, fmt.Errorf("defn: Wrong number of arguments")
	}

	s, ok := args[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("defn: First argument must be a symbol")
	}

	if fn, err := BuiltInFn(e, args[1:]); err != nil {
		return nil, err
	} else {
		e.SetVar(s.Name, fn)
		return fn, nil
	}
}

func BuiltInDefmacro(e Environment, args []Form) (Form, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("defmacro: Wrong number of arguments")
	}

	s, ok := args[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("defmacro: First argument must be a symbol")
	}

	fn, err := BuiltInFn(e, args[1:])

	if err != nil {
		return nil, err
	}

	m := Macro{
		Args: fn.(UserFunction).Args,
		Body: fn.(UserFunction).Body,
	}
	e.SetVar(s.Name, m)
	return m, nil
}

func BuiltInMacroExpand(e Environment, args []Form) (Form, error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("macroexpand: Wrong number of arguments")
	}

	var mname Form
	var err error
	if mname, err = args[0].Eval(e); err != nil {
		return nil, err
	}

	m, ok := mname.(Macro)
	if !ok {
		return nil, fmt.Errorf("macroexpand: First argument must evaluate to a macro")
	}

	return m.Expand(args[1:])
}

func BuiltInFn(e Environment, args []Form) (Form, error) {
	if len(args) != 2 {
		return nil, fmt.Errorf("fn: Wrong number of arguments")
	}

	fn := UserFunction{}

	if fargs, ok := args[0].(*Vector); !ok {
		return nil, fmt.Errorf("fn: first argument must be a vector")
	} else {
		if symbols, ok := castToSymbols(fargs.Items); !ok {
			return nil, fmt.Errorf("fn: all arguments must evaluate to symbols")
		} else {
			fn.Args = symbols
		}
	}

	if fbody, ok := args[1].(*List); !ok {
		return nil, fmt.Errorf("fn: second argument must be a list")
	} else {
		fn.Body = fbody
	}

	return fn, nil
}

func castToSymbols(a []Form) ([]Symbol, bool) {
	res := make([]Symbol, len(a))
	for idx, f := range a {
		if s, ok := f.(Symbol); ok {
			res[idx] = s
		} else {
			return nil, false
		}
	}
	return res, true
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