package main

import (
	"fmt"
	"strings"
)

type Vector struct {
	Items []Form
}

func (l *Vector) String() string {
	var s []string

	for _, i := range l.Items {
		s = append(s, fmt.Sprint(i))
	}

	return "[" + strings.Join(s, " ") + "]"
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
