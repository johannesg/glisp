package main

import ()

type CoreFunctions struct {
	Name string
}

func (c *CoreFunctions) Eval(e Environment) (Form, error) {
	return c, nil
}

func (c *CoreFunctions) List(args []Form) Form {
	return &List{Items: args}
}

func (c *CoreFunctions) First(l *List) Form {
	return l.First()
}

func (c *CoreFunctions) Rest(l *List) Form {
	return l.Rest()
}

func (c *CoreFunctions) Cons(v Form, l *List) Form {
	items := make([]Form, len(l.Items)+1)
	items[0] = v
	for i, v := range l.Items {
		items[i+1] = v
	}
	return &List{
		Items: items,
	}
}

func (c *CoreFunctions) Conj(l *List, v Form) Form {
	items := make([]Form, len(l.Items)+1)
	items[0] = v
	for i, v := range l.Items {
		items[i+1] = v
	}
	return &List{
		Items: items,
	}
}
