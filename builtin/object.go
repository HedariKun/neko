package neko

type Method func([]Object) Object

type Object interface {
	SetField(string, Object)
	GetField(string) Object
	CallMethod(string, []Object) Object
	GetMethod(string) Method
}
