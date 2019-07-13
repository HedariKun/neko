package neko

import (
	builtin "github.com/hedarikun/neko/builtin"
)

type Fun func(args []builtin.Object) builtin.Object

type ScopeInterface interface {
	GetVariable(string) builtin.Object
	SetVariable(string, builtin.Object)
	RegisterFun(string, func([]builtin.Object) builtin.Object)
	ExecuteFun(string, []builtin.Object) builtin.Object
	GetFun(string) builtin.Object
}

type Scope struct {
	Functions map[string]builtin.FunObject
	Variables map[string]builtin.Object
	Outer     ScopeInterface
}

func (s *Scope) GetVariable(name string) builtin.Object {
	if _, ok := s.Functions[name]; !ok {
		return s.Variables[name]
	}

	p, ok := s.Outer.(*Scope)
	if !ok {
		return s.Outer.GetVariable(name)
	}
	return p.GetVariable(name)
}

func (s *Scope) SetVariable(name string, value builtin.Object) {
	s.Variables[name] = value
}

func (s *Scope) RegisterFun(name string, fun func([]builtin.Object) builtin.Object) {
	s.Functions[name] = builtin.NewFun(fun)
}

func (s *Scope) ExecuteFun(name string, args []builtin.Object) builtin.Object {
	return s.Functions[name].CallMethod("call", args)
}

func (s *Scope) GetFun(name string) builtin.Object {
	if _, ok := s.Functions[name]; ok {
		return s.Functions[name]
	}

	p, ok := s.Outer.(*Scope)
	if !ok {
		return s.Outer.GetFun(name)
	}
	return p.GetFun(name)
}

func NewScope() *Scope {
	s := Scope{}
	s.Functions = make(map[string]builtin.FunObject, 0)
	s.Variables = make(map[string]builtin.Object, 0)
	return &s
}

type Global struct {
	Functions map[string]builtin.FunObject
	Variables map[string]builtin.Object
}

func (g *Global) GetVariable(name string) builtin.Object {
	return g.Variables[name]
}

func (g *Global) SetVariable(name string, value builtin.Object) {
	g.Variables[name] = value
}

func (g *Global) RegisterFun(name string, fun func([]builtin.Object) builtin.Object) {
	g.Functions[name] = builtin.NewFun(fun)
}

func (g *Global) ExecuteFun(name string, args []builtin.Object) builtin.Object {
	return g.Functions[name].CallMethod("call", args)
}

func (g *Global) GetFun(name string) builtin.Object {
	if _, ok := g.Functions[name]; !ok {
		return nil
	}
	return g.Functions[name]
}
