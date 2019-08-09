package neko

type StructObject struct {
	EmptyObject
	Props []string
	Meth  map[string]Method
}

func NewStruct(props []string) Object {
	so := &StructObject{
		Props: props,
	}
	so.Fields = make(map[string]Object, 0)
	so.Methods = make(map[string]Method, 0)
	so.Meth = make(map[string]Method, 0)

	so.Methods["new"] = func(args []Object) Object {
		no := NewObject()
		for index, value := range so.Props {
			if len(args) > index {
				no.SetField(value, args[index])
			} else {
				no.SetField(value, NewString(""))
			}
		}
		return no
	}

	return so
}
