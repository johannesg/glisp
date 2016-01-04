package main

import (
	"fmt"
	"reflect"
)

type Interop struct {
	Name string
}

func (i Interop) Eval(e Environment) (Form, error) {
	return i, nil
}

func (i Interop) Invoke(e Environment, args []Form) (ret Form, err error) {
	if len(args) < 1 {
		return nil, fmt.Errorf("Wrong number of arguments")
	}

	var instance Form
	if instance, err = args[0].Eval(e); err != nil {
		return
	}

	t := reflect.TypeOf(instance)

	var m reflect.Method
	var ok bool
	if m, ok = t.MethodByName(i.Name); !ok {
		return nil, fmt.Errorf("Method %v on type %v not found", i.Name, t.Name())
	}

	params := make([]reflect.Value, len(args))
	params[0] = reflect.ValueOf(instance)
	for idx, a := range args[1:] {
		params[idx] = reflect.ValueOf(a)
	}

	res := m.Func.Call(params)
	if len(res) == 0 {
		return nil, nil
	}

	if len(res) >= 1 {
		if res[0].IsNil() {
			ret = nil
		} else {
			ret = res[0].Interface().(Form)
		}
	}

	if len(res) > 1 {
		if res[1].IsNil() {
			err = nil
		} else {
			err = res[1].Interface().(error)
		}
	}

	if len(res) > 2 {
		return nil, fmt.Errorf("Not implemented")
	}
	return
}
