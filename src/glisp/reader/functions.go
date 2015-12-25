package reader

import (
	"fmt"
)

type LispFunction func([]Form) (Form, error)

var builtIns = map[string]LispFunction{
	"add":  BuiltInAdd,
	"eval": BuiltInEval,
}

func InvokeBuiltIn(symbol Form, args []Form) (Form, error) {
	s, ok := symbol.(Symbol)
	if !ok {
		return nil, fmt.Errorf("Expected symbol")
	}

	f, ok := builtIns[s.Name]
	if !ok {
		return nil, fmt.Errorf("Function not found")
	}

	return f(args)
}

func BuiltInAdd(args []Form) (Form, error) {
	result := 0
	for _, i := range args {
		f, err := i.Eval()
		if err != nil {
			return nil, err
		}

		n, ok := f.(Number)
		if !ok {
			return nil, fmt.Errorf("Expected number")
		}

		result += n.Val
	}
	return Number{Val: result}, nil
}

func BuiltInEval(args []Form) (Form, error) {
	q, ok := args[0].(*QForm)
	if !ok {
		return args[0].Eval()
	}

	return q.Form.Eval()
}
