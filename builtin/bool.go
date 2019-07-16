package neko

type BoolObject struct {
	Value   bool
	Mut     bool
	Fields  map[string]Object
	Methods map[string]Method
}

func (bo *BoolObject) IsMutable() bool {
	return bo.Mut
}

func (bo *BoolObject) SetMutable(value bool) {
	bo.Mut = value
}

func (bo *BoolObject) SetField(name string, val Object) {
	bo.Fields[name] = val
}

func (bo *BoolObject) GetField(name string) Object {
	return bo.Fields[name]
}

func (bo *BoolObject) CallMethod(name string, args []Object) Object {
	return bo.Methods[name](args)
}

func (bo *BoolObject) GetMethod(name string) Method {
	return bo.Methods[name]
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
