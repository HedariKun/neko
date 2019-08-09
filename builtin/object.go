package neko

type Method func([]Object) Object

type Object interface {
	SetField(string, Object)
	GetField(string) Object
	CallMethod(string, []Object) Object
	SetMethod(string, Method)
	GetMethod(string) Method
	IsMutable() bool
	SetMutable(bool)
}

type EmptyObject struct {
	Mut     bool
	Fields  map[string]Object
	Methods map[string]Method
}

func (eo *EmptyObject) SetField(name string, value Object) {
	eo.Fields[name] = value
}

func (eo *EmptyObject) GetField(name string) Object {
	return eo.Fields[name]
}

func (eo *EmptyObject) CallMethod(name string, args []Object) Object {
	return eo.Methods[name](args)
}

func (eo *EmptyObject) SetMethod(name string, method Method) {
	eo.Methods[name] = method
}

func (eo *EmptyObject) GetMethod(name string) Method {
	return eo.Methods[name]
}

func (eo *EmptyObject) IsMutable() bool {
	return eo.Mut
}

func (eo *EmptyObject) SetMutable(value bool) {
	eo.Mut = value
}

func NewObject() Object {
	eo := &EmptyObject{}
	eo.Fields = make(map[string]Object, 0)
	eo.Methods = make(map[string]Method, 0)
	return eo
}
