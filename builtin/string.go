package neko

type StringObject struct {
	Value   string
	Fields  map[string]Object
	Methods map[string]Method
}

func (so StringObject) SetField(name string, val Object) {
	so.Fields[name] = val
}

func (so StringObject) GetField(name string) Object {
	return so.Fields[name]
}

func (so StringObject) CallMethod(name string, args []Object) Object {
	return so.Methods[name](args)
}

func (so StringObject) GetMethod(name string) Method {
	return so.Methods[name]
}

func NewString(value string) StringObject {
	so := StringObject{
		Value: value,
	}
	so.Methods = make(map[string]Method, 0)
	so.Methods["toString"] = func(args []Object) Object {
		return NewString(so.Value)
	}
	so.Methods["add"] = func(args []Object) Object {
		arg := args[0]
		right, _ := arg.CallMethod("toString", nil).(StringObject)
		return NewString(so.Value + right.Value)
	}
	so.Methods["indexOf"] = func(args []Object) Object {
		val, ok := args[0].(NumberObject)
		if !ok {
			// error handling
		}
		if len(so.Value) < int(val.Value) {
			// error
		}
		return NewString(string([]rune(so.Value)[int(val.Value)]))
	}
	return so
}
