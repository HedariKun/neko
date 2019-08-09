package neko

type ArrayObject struct {
	EmptyObject
	Values []Object
}

func (ao *ArrayObject) addElement(element Object) {
	if !ao.IsMutable() {
		// error handling
	}
	ao.Values = append(ao.Values, element)
}

func NewArray(values []Object) Object {
	arr := &ArrayObject{
		Values: values,
	}

	arr.Fields = make(map[string]Object, 0)
	arr.Methods = make(map[string]Method, 0)

	arr.Methods["push"] = func(args []Object) Object {
		if args[0] == nil {
			// error handling
		}
		for _, element := range args {
			arr.addElement(element)
		}
		return arr
	}

	arr.Methods["indexOf"] = func(args []Object) Object {
		val, ok := args[0].(*NumberObject)
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
