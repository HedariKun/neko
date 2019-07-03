package neko

type Method func([]Object) Object

type Object interface {
	SetField(string, Object)
	GetField(string) Object
	CallMethod(string, []Object) Object
}

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
	return NumberObject{
		Value: value,
	}
}
