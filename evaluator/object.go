package neko

import (
	"strconv"
)

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
	no := NumberObject{
		Value: value,
	}
	no.Methods = make(map[string]Method, 0)
	no.Methods["toString"] = func(args []Object) Object {
		val := strconv.FormatFloat(no.Value, 'f', -1, 64)
		return NewString(val)
	}
	return no
}

type StringObject struct {
	Value   string
	Fields  map[string]Object
	Methods map[string]Method
}

func (so StringObject) SetField(name string, val Object) {
	so.Fields[name] = val
}

func (so StringObject) GetField(name string) Object {
	return so.Fields[name]
}

func (so StringObject) CallMethod(name string, args []Object) Object {
	return so.Methods[name](args)
}

func NewString(value string) StringObject {
	so := StringObject{
		Value: value,
	}
	so.Methods = make(map[string]Method, 0)
	so.Methods["toString"] = func(args []Object) Object {
		return NewString(so.Value)
	}
	return so
}
