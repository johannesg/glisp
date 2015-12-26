package main

import (
	"fmt"
)

func (n Number) Eval() (Form, error) {
	return n, nil
}

func (s Symbol) Eval() (Form, error) {
	return s, nil
}

func (l Literal) Eval() (Form, error) {
	return l, nil
}

func (l *List) Eval() (Form, error) {
	if len(l.Items) == 0 {
		return l, nil
	}

	s, ok := l.Items[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("Expected symbol")
	}

	return InvokeBuiltIn(s, l.Items[1:])
}

func (q *QForm) Eval() (Form, error) {
	return q, nil
}

func (b Boolean) Eval() (Form, error) {
	return b, nil
}
