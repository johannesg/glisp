package main

import (
	"fmt"
	"strings"
)

type Form interface {
	Eval(Environment) (Form, error)
}

type List struct {
	Items []Form
}

type Vector struct {
	Items []Form
}

type Symbol struct {
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

func (n Number) String() string {
	return fmt.Sprint(n.Val)
}

func (s Symbol) String() string {
	return fmt.Sprint(s.Name)
}

func (l Literal) String() string {
	return "\"" + l.Val + "\""
}

func (l *List) String() string {
	var s []string

	for _, i := range l.Items {
		s = append(s, fmt.Sprint(i))
	}

	return "(" + strings.Join(s, " ") + ")"
}

func (l *Vector) String() string {
	var s []string

	for _, i := range l.Items {
		s = append(s, fmt.Sprint(i))
	}

	return "[" + strings.Join(s, " ") + "]"
}

func (q *QForm) String() string {
	return fmt.Sprintf("'%v", q.Form)
}
