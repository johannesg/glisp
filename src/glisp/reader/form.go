package reader

import "fmt"

type Form interface {
	Eval() (Form, error)
}

type List struct {
	items []Form
}

type Symbol struct {
	name string
}

type Literal struct {
	val string
}

type Number struct {
	val int
}

func (n Number) String() string {
	return fmt.Sprint(n.val)
}

func (s Symbol) String() string {
	return fmt.Sprint(s.name)
}

func (l Literal) String() string {
	return fmt.Sprint(l.val)
}

func (l *List) String() string {
	var s string
	s += "("
	for _, i := range l.items {
		s += fmt.Sprint(i)
		s += " "
	}

	s = ")"
	return s
}
