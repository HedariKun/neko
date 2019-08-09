package neko

type FunObject struct {
	EmptyObject
	body Method
}

func NewFun(body Method) Object {
	fo := &FunObject{
		body: body,
	}

	fo.Methods = make(map[string]Method, 0)
	fo.Fields = make(map[string]Object, 0)

	fo.Methods["call"] = func(args []Object) Object {
		return fo.body(args)
	}

	return fo
}
