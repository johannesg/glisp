package main

import (
	"io/ioutil"
)

type Environment interface {
	Load(string) (Form, error)
	Eval(string) (Form, error)
	SetVar(string, Form)
	Var(string) (Form, bool)
	Vars() SymbolTable
}

type SymbolTable map[string]Form

type environment struct {
	parent Environment
	vars   SymbolTable
}

var builtInVars = map[string]Form{
	"::core": &CoreFunctions{
		Name: "The Core!",
	},
}

func NewEnvironment(parent Environment) Environment {
	return &environment{
		parent: parent,
		vars:   make(SymbolTable)}
}

func (e *environment) Load(filename string) (Form, error) {
	buf, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return e.Eval(string(buf))
}

func (e *environment) Eval(input string) (ret Form, err error) {
	r := NewReader(input)
	for {
		var f Form
		f, err = r.Read()
		if f == nil || err != nil {
			return
		}

		ret, err = f.Eval(e)

		if err != nil {
			return
		}
	}
}

func (e *environment) SetVar(name string, f Form) {
	e.vars[name] = f
}

func (e *environment) Var(name string) (v Form, ok bool) {
	if v, ok = builtInVars[name]; ok {
		return
	}

	if v, ok = e.vars[name]; ok {
		return
	}

	if e.parent != nil {
		return e.parent.Var(name)
	}

	v, ok = builtIns[name]
	return
}

func (e *environment) Vars() SymbolTable {
	return e.vars
}
