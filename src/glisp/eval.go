package main

import (
	"fmt"
)

func (n Number) Eval(e Environment) (Form, error) {
	return n, nil
}

func (s Symbol) Eval(e Environment) (Form, error) {
	v, ok := e.Var(s.Name)
	if ok {
		return v, nil
	}
	return s, nil
}

func (l Literal) Eval(e Environment) (Form, error) {
	return l, nil
}

func (l *List) Eval(e Environment) (Form, error) {
	if len(l.Items) == 0 {
		return l, nil
	}

	items := make([]Form, len(l.Items))

	for idx, i := range l.Items {
		ret, err := i.Eval(e)
		if err != nil {
			return nil, err
		}

		items[idx] = ret
	}

	f, ok := items[0].(Function)
	if !ok {
		return nil, fmt.Errorf("First argument must evaluate to a function")
	}

	return f.Invoke(e, items[1:])
}

func (v *Vector) Eval(e Environment) (Form, error) {
	if len(v.Items) == 0 {
		return v, nil
	}

	items := make([]Form, len(v.Items))
	for idx, i := range v.Items {

		ret, err := i.Eval(e)
		if err != nil {
			return nil, err
		}

		items[idx] = ret
	}

	return &Vector{Items: items}, nil
}

func (q *QForm) Eval(e Environment) (Form, error) {
	return q.Form, nil
}

func (b Boolean) Eval(e Environment) (Form, error) {
	return b, nil
}

func (f BuiltInFunction) Eval(e Environment) (Form, error) {
	return f, nil
}

func (f UserFunction) Eval(e Environment) (Form, error) {
	return f, nil
}
