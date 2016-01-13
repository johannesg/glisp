package main

import (
	"fmt"
	"strings"
)

type List struct {
	Items []Form
}

func (l *List) String() string {
	var s []string

	for _, i := range l.Items {
		s = append(s, fmt.Sprint(i))
	}

	return "(" + strings.Join(s, " ") + ")"
}

func (l *List) Eval(e Environment) (Form, error) {
	if len(l.Items) == 0 {
		return l, nil
	}

	var fname Form
	var err error
	if fname, err = l.Items[0].Eval(e); err != nil {
		return nil, err
	}

	f, ok := fname.(Callable)
	if !ok {
		return nil, fmt.Errorf("First argument must evaluate to a function")
	}

	return f.Call(e, l.Items[1:])
}

func (l *List) Expand(argmap map[string]Form) (Form, error) {
	a := make([]Form, len(l.Items))
	var err error
	for idx, i := range l.Items {
		var res Form
		if exp, ok := i.(Expandable); ok {
			res, err = exp.Expand(argmap)
		} else if sym, ok := i.(Symbol); ok {
			if res, ok = argmap[sym.Name]; !ok {
				res = i
			}
		} else {
			res = i
		}

		if err != nil {
			return nil, err
		}
		a[idx] = res
	}
	return &List{
		Items: a,
	}, err
}

func (l *List) First() Form {
	return l.Items[0]
}

func (l *List) Rest() Form {
	return &List{Items: l.Items[1:]}
}
