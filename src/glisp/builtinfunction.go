package main

import (
	"fmt"
	"log"
)

type BuiltInFunction func(BuiltInFunctions, Environment, []Form) (Form, error)

type BuiltInFunctions struct{}

func (f BuiltInFunction) Eval(e Environment) (Form, error) {
	return f, nil
}

func (f BuiltInFunction) Call(e Environment, args []Form) (Form, error) {
	return f(BuiltInFunctions{}, e, args)
}

var builtIns = map[string]BuiltInFunction{
	"do":          BuiltInFunctions.Do,
	"def":         BuiltInFunctions.Def,
	"defn":        BuiltInFunctions.Defn,
	"defmacro":    BuiltInFunctions.Defmacro,
	"fn":          BuiltInFunctions.Fn,
	"var":         BuiltInFunctions.Var,
	"vars":        BuiltInFunctions.Vars,
	"macroexpand": BuiltInFunctions.MacroExpand,
}

func (BuiltInFunctions) Add(e Environment, args []Form) (Form, error) {
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

func (BuiltInFunctions) Do(e Environment, args []Form) (ret Form, err error) {
	for _, a := range args {
		ret, err = a.Eval(e)
		if err != nil {
			return
		}

		ret, err = ret.Eval(e)
		if err != nil {
			return
		}
	}
	return
}

func (BuiltInFunctions) Def(e Environment, args []Form) (Form, error) {
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

func (f BuiltInFunctions) Defn(e Environment, args []Form) (Form, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("defn: Wrong number of arguments")
	}

	s, ok := args[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("defn: First argument must be a symbol")
	}

	if fn, err := f.Fn(e, args[1:]); err != nil {
		return nil, err
	} else {
		e.SetVar(s.Name, fn)
		return fn, nil
	}
}

func (f BuiltInFunctions) Defmacro(e Environment, args []Form) (Form, error) {
	if len(args) < 3 {
		return nil, fmt.Errorf("defmacro: Wrong number of arguments")
	}

	s, ok := args[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("defmacro: First argument must be a symbol")
	}

	fn, err := f.Fn(e, args[1:])

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

func (BuiltInFunctions) MacroExpand(e Environment, args []Form) (Form, error) {
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

func (BuiltInFunctions) Fn(e Environment, args []Form) (Form, error) {
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

func (BuiltInFunctions) Vars(e Environment, args []Form) (Form, error) {
	// items := []Form{}

	for k, v := range e.Vars() {
		log.Printf("%v: %#v", k, v)
	}

	return nil, nil
}

func (BuiltInFunctions) Var(e Environment, args []Form) (Form, error) {
	// items := []Form{}

	if len(args) != 1 {
		return nil, fmt.Errorf("Wrong number of arguments")
	}

	var name string
	if s, ok := args[0].(Symbol); ok {
		name = s.Name
	}

	if k, ok := args[0].(Keyword); ok {
		name = k.Name
	}

	if ret, ok := e.Var(name); ok {
		return ret, nil
	}

	return nil, fmt.Errorf("Var not defined")
}

// func BuiltInEquality(args []Form) (Form, error) {
// 	if len(args) == 0 {
// 		return nil, fmt.Errorf("Wrong number of arguments")
// 	}

// 	for a := range args {

// 	}
// }
