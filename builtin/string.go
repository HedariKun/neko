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

func NewString(value string) StringObject {
	so := StringObject{
		Value: value,
	}
	so.Methods = make(map[string]Method, 0)
	so.Methods["toString"] = func(args []Object) Object {
		return NewString(so.Value)
	}
	so.Methods["add"] = func(args []Object) Object {
		right, _ := args[0].(StringObject)
		return NewString(so.Value + right.Value)
	}
	return so
}
