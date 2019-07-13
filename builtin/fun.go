package neko

type FunObject struct {
	body    Method
	Fields  map[string]Object
	Methods map[string]Method
}

func (fo FunObject) SetField(name string, val Object) {
	fo.Fields[name] = val
}

func (fo FunObject) GetField(name string) Object {
	return fo.Fields[name]
}

func (fo FunObject) CallMethod(name string, args []Object) Object {
	return fo.Methods[name](args)
}

func (fo FunObject) GetMethod(name string) Method {
	return fo.Methods[name]
}

func NewFun(body Method) FunObject {
	fo := FunObject{
		body: body,
	}

	fo.Methods = make(map[string]Method, 0)
	fo.Fields = make(map[string]Object, 0)

	fo.Methods["call"] = func(args []Object) Object {
		return fo.body(args)
	}

	return fo
}
