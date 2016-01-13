package main

import (
	"fmt"
)

type Symbol struct {
	Name string
}

type Keyword struct {
	Name string
}

type Literal struct {
	Val string
}

type Number struct {
	Val int
}

type QForm struct {
	Form Form
}

type Boolean bool

func (s Symbol) String() string {
	return s.Name
}

func (s Keyword) String() string {
	return s.Name
}

func (l Literal) String() string {
	return "\"" + l.Val + "\""
}

func (n Number) String() string {
	return fmt.Sprint(n.Val)
}

func (q *QForm) String() string {
	return fmt.Sprintf("'%v", q.Form)
}

func (s Symbol) Eval(e Environment) (Form, error) {
	if v, ok := e.Var(s.Name); ok {
		return v, nil
	}

	return nil, fmt.Errorf("Unable to resolve symbol %v", s.Name)
}

func (k Keyword) Eval(e Environment) (Form, error) {
	return k, nil
}

func (l Literal) Eval(e Environment) (Form, error) {
	return l, nil
}

func (n Number) Eval(e Environment) (Form, error) {
	return n, nil
}

func (q *QForm) Eval(e Environment) (Form, error) {
	return q.Form, nil
}

func (b Boolean) Eval(e Environment) (Form, error) {
	return b, nil
}
