package main

type Environment interface {
	Eval(string) (Form, error)
	Invoke(Symbol, []Form) (Form, error)
	SetVar(string, Form)
	Var(string) (Form, bool)
}

type SymbolTable map[string]Form

type environment struct {
	vars SymbolTable
}

func NewEnvironment() Environment {
	return &environment{vars: make(SymbolTable)}
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

func (e *environment) Var(name string) (Form, bool) {
	v, ok := e.vars[name]
	return v, ok
}
