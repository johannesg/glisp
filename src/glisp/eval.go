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

	s, ok := l.Items[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("Expected symbol")
	}

	var args []Form
	for _, i := range l.Items[1:] {
		ret, err := i.Eval(e)
		if err != nil {
			return nil, err
		}

		// log.Printf("List item: %v", ret)

		args = append(args, ret)
	}

	return e.Invoke(s, args)
}

func (q *QForm) Eval(e Environment) (Form, error) {
	return q, nil
}

func (b Boolean) Eval(e Environment) (Form, error) {
	return b, nil
}
