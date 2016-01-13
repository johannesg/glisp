package main

import (
	"fmt"
)

type ArgList interface {
	Set(args []Form, pos int, e Environment) Boolean
}

type ArgVector struct {
	Args []ArgList
}

type ArgItem struct {
	Name string
}

type ArgRest struct {
	Pos  int
	Name string
}

type ArgAll struct {
	Name string
}

func NewArgList(args []Form) (*ArgVector, error) {
	var items []ArgList
	restIndex := -1

	for i := 0; i < len(args); i++ {
		if s, ok := args[i].(Symbol); ok {
			if s.Name == "&" {
				if restIndex == -1 {
					restIndex = i
				}

				i++
				if i == len(args) {
					return nil, fmt.Errorf("Symbol must follow after &")
				}
				if s, ok := args[i].(Symbol); ok {
					items = append(items, &ArgRest{Pos: restIndex, Name: s.Name})
				} else {
					return nil, fmt.Errorf("Symbol must follow after &")
				}

			} else {
				if restIndex != -1 {
					return nil, fmt.Errorf("Symbols after & not allowed")
				}

				items = append(items, &ArgItem{Name: s.Name})
			}
		} else if k, ok := args[i].(Keyword); ok && i == len(args)-1 {
			if k.Name != ":as" {
				return nil, fmt.Errorf("The only keyword allowed is :as and it must come last")
			}

			items = append(items, &ArgAll{Name: k.Name})
		} else if v, ok := args[i].(*Vector); ok {

			argVector, err := NewArgList(v.Items)
			if err != nil {
				return nil, err
			}
			items = append(items, argVector)
		} else {
			return nil, fmt.Errorf("Invalid param, %v", args[i])
		}
	}

	return &ArgVector{Args: items}, nil
}

func (a *ArgVector) Match(args []Form, e Environment) Boolean {
	for idx, i := range a.Args {
		if !i.Set(args, idx, e) {
			return false
		}
	}

	return true
}

func (a *ArgVector) Set(args []Form, pos int, e Environment) Boolean {
	v, ok := args[pos].(*Vector)
	if !ok {
		return false
	}

	return a.Match(v.Items, e)
}

func (a *ArgItem) Set(args []Form, pos int, e Environment) Boolean {
	if pos >= len(args) {
		return false
	}

	e.SetVar(a.Name, args[pos])
	return true
}

func (a *ArgRest) Set(args []Form, pos int, e Environment) Boolean {
	if pos >= len(args) {
		e.SetVar(a.Name, nil)
	} else {
		e.SetVar(a.Name, &Vector{Items: args[pos:]})
	}
	return true
}

func (a *ArgAll) Set(args []Form, pos int, e Environment) Boolean {
	e.SetVar(a.Name, &Vector{Items: args})
	return true
}

type UserFunction struct {
	Args *ArgVector
	Body *List
}

func (f UserFunction) Eval(e Environment) (Form, error) {
	return f, nil
}

func (f UserFunction) Call(e Environment, args []Form) (Form, error) {
	eargs := make([]Form, len(args))

	for idx, a := range args {
		if ea, err := a.Eval(e); err != nil {
			return nil, err
		} else {
			eargs[idx] = ea
		}
	}

	local := NewEnvironment(e)
	if !f.Args.Match(eargs, local) {
		return nil, fmt.Errorf("Could not match arguments")
	}

	// for idx, a := range args {
	// 	if da, err := a.Eval(e); err == nil {
	// 		local.SetVar(f.Args[idx].Name, da)
	// 	} else {
	// 		return nil, err
	// 	}
	// }

	return f.Body.Eval(local)
}

// func (f UserFunction) MatchArgs(e Environment, args []Form) {
// 	var rest []Form
// 	insideRest := false
// 	for idx, a := range f.Args {
// 		if a.Name == "&" && rest != nil {
// 			rest = args[idx:]
// 		}
// 	}

// }
