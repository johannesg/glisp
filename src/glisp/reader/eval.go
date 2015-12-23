package reader

import "fmt"

func (n Number) Eval() (Form, error) {
	return n, nil
}

func (s Symbol) Eval() (Form, error) {
	return s, nil
}

func (l Literal) Eval() (Form, error) {
	return l, nil
}

func (l List) Eval() (Form, error) {
	if len(l.items) == 0 {
		return l, nil
	}

	s, ok := l.items[0].(Symbol)
	if !ok {
		return nil, fmt.Errorf("Expected symbol")
	}

	switch s.name {
	case "add":
		return Add(l.items[1:])
	default:
		return nil, fmt.Errorf("Not implemented: %v", s.name)
	}
}

func Add(items []Form) (Form, error) {
	result := 0
	for _, i := range items {
		f, err := i.Eval()
		if err != nil {
			return nil, err
		}

		n, ok := f.(Number)
		if !ok {
			return nil, fmt.Errorf("Expected number")
		}

		result += n.val
	}
	return Number{val: result}, nil
}
