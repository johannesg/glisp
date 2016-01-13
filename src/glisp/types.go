package main

type Form interface {
	Eval(Environment) (Form, error)
}

type Callable interface {
	Call(Environment, []Form) (Form, error)
}

type Expandable interface {
	Expand(map[string]Form) (Form, error)
}
