package reader

import "fmt"

type Form interface {
	Eval() (Form, error)
}

type List struct {
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
	var s string
	s += "("
	for _, i := range l.Items {
		s += fmt.Sprint(i)
		s += " "
	}

	s += ")"
	return s
}

func (q *QForm) String() string {
	return fmt.Sprintf("'%v", q.Form)
}
