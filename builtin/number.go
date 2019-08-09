package neko

import (
	"strconv"
)

type NumberObject struct {
	EmptyObject
	Value float64
}

func NewNumber(value float64) Object {
	no := &NumberObject{
		Value: value,
	}
	no.Fields = make(map[string]Object, 0)
	no.Methods = make(map[string]Method, 0)
	no.Methods["toString"] = func(args []Object) Object {
		val := strconv.FormatFloat(no.Value, 'f', -1, 64)
		return NewString(val)
	}
	no.Methods["toValue"] = func(args []Object) Object {
		return no
	}
	no.Methods["add"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].CallMethod("toValue", nil).(*NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value + right.Value)
	}

	no.Methods["subtract"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value - right.Value)
	}

	no.Methods["multiply"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value * right.Value)
	}

	no.Methods["divide"] = func(args []Object) Object {
		if len(args) <= 0 {
			// error
		}
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewNumber(no.Value / right.Value)
	}

	no.Methods["equal"] = func(args []Object) Object {
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value == right.Value)
	}

	no.Methods["greater"] = func(args []Object) Object {
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value > right.Value)
	}

	no.Methods["greaterOrEqual"] = func(args []Object) Object {
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value >= right.Value)
	}

	no.Methods["lower"] = func(args []Object) Object {
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value < right.Value)
	}

	no.Methods["lowerOrEqual"] = func(args []Object) Object {
		right, ok := args[0].(*NumberObject)
		if !ok {
			// error
		}
		return NewBool(no.Value <= right.Value)
	}

	return no
}
