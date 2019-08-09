package neko

type BoolObject struct {
	EmptyObject
	Value bool
}

func NewBool(value bool) Object {
	bo := &BoolObject{
		Value: value,
	}
	bo.Methods = make(map[string]Method, 0)
	bo.Methods["toString"] = func(args []Object) Object {
		if bo.Value == true {
			return NewString("true")
		} else {
			return NewString("false")
		}
	}
	bo.Methods["toValue"] = func(args []Object) Object {
		if bo.Value {
			return NewNumber(1)
		} else {
			return NewNumber(0)
		}
	}
	bo.Methods["or"] = func(args []Object) Object {
		val, ok := args[0].(*BoolObject)
		if !ok {
			// error handling
		}
		return NewBool(bo.Value || val.Value)
	}
	bo.Methods["and"] = func(args []Object) Object {
		val, ok := args[0].(*BoolObject)
		if !ok {
			// error handling
		}
		return NewBool(bo.Value && val.Value)
	}
	return bo
}
