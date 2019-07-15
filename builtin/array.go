package neko

type ArrayObject struct {
	Values  []Object
	Fields  map[string]Object
	Methods map[string]Method
}

func (ao ArrayObject) SetField(name string, val Object) {
	ao.Fields[name] = val
}

func (ao ArrayObject) GetField(name string) Object {
	return ao.Fields[name]
}

func (ao ArrayObject) CallMethod(name string, args []Object) Object {
	return ao.Methods[name](args)
}

func (ao ArrayObject) GetMethod(name string) Method {
	return ao.Methods[name]
}

func NewArray(values []Object) Object {
	arr := ArrayObject{
		Values:  values,
		Fields:  make(map[string]Object, 0),
		Methods: make(map[string]Method, 0),
	}

	arr.Methods["indexOf"] = func(args []Object) Object {
		val, ok := args[0].(NumberObject)
		if !ok {
			// error handling
		}
		if len(arr.Values) < int(val.Value) {
			// error handling
		}
		return arr.Values[int(val.Value)]
	}
	return arr
}
