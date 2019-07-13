package neko

type BoolObject struct {
	Value   bool
	Fields  map[string]Object
	Methods map[string]Method
}

func (bo BoolObject) SetField(name string, val Object) {
	bo.Fields[name] = val
}

func (bo BoolObject) GetField(name string) Object {
	return bo.Fields[name]
}

func (bo BoolObject) CallMethod(name string, args []Object) Object {
	return bo.Methods[name](args)
}

func (bo BoolObject) GetMethod(name string) Method {
	return bo.Methods[name]
}

func NewBool(value bool) BoolObject {
	bo := BoolObject{
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
	return bo
}
