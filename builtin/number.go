package neko

import (
	"strconv"
)

type NumberObject struct {
	Value   float64
	Fields  map[string]Object
	Methods map[string]Method
}

func (no NumberObject) SetField(name string, val Object) {
	no.Fields[name] = val
}

func (no NumberObject) GetField(name string) Object {
	return no.Fields[name]
}

func (no NumberObject) CallMethod(name string, args []Object) Object {
	return no.Methods[name](args)
}

func NewNumber(value float64) NumberObject {
	no := NumberObject{
		Value: value,
	}
	no.Methods = make(map[string]Method, 0)
	no.Methods["toString"] = func(args []Object) Object {
		val := strconv.FormatFloat(no.Value, 'f', -1, 64)
		return NewString(val)
	}

	no.Methods["add"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value + right.Value)
	}

	no.Methods["subtract"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value - right.Value)
	}

	no.Methods["multiply"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value * right.Value)
	}

	no.Methods["divide"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value / right.Value)
	}

	no.Methods["equal"] = func(args []Object) Object {
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value == right.Value)
	}

	no.Methods["greater"] = func(args []Object) Object {
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value > right.Value)
	}

	no.Methods["greaterOrEqual"] = func(args []Object) Object {
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value >= right.Value)
	}

	no.Methods["lower"] = func(args []Object) Object {
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value < right.Value)
	}

	no.Methods["lowerOrEqual"] = func(args []Object) Object {
		right, ok := args[0].(NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value <= right.Value)
	}

	return no
}
