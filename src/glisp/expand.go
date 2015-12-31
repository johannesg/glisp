package main

func (m Macro) Expand(args []Form) (Form, error) {
	argmap := make(map[string]Form)

	for idx, a := range args {
		argmap[m.Args[idx].Name] = a
	}

	return m.Body.Expand(argmap)
}

type Expandable interface {
	Expand(map[string]Form) (Form, error)
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
