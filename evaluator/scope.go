package neko

import (
	builtin "github.com/hedarikun/neko/builtin"
)

type Fun func([]builtin.Object) builtin.Object

type ScopeInterface interface {
	GetVariable(string) builtin.Object
	SetVariable(string, builtin.Object)
	RegisterFun(string, Fun)
	ExecuteFun(string, []builtin.Object) builtin.Object
	GetFun(string) Fun
}

type Scope struct {
	Functions map[string]Fun
	Variables map[string]builtin.Object
	Outer     ScopeInterface
}

func (s *Scope) GetVariable(name string) builtin.Object {
	if s.Functions[name] != nil {
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

func (s *Scope) RegisterFun(name string, fun Fun) {
	s.Functions[name] = fun
}

func (s *Scope) ExecuteFun(name string, args []builtin.Object) builtin.Object {
	return s.Functions[name](args)
}

func (s *Scope) GetFun(name string) Fun {
	if s.Functions[name] != nil {
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
	s.Functions = make(map[string]Fun, 0)
	s.Variables = make(map[string]builtin.Object, 0)
	return &s
}

type Global struct {
	Functions map[string]Fun
	Variables map[string]builtin.Object
}

func (g *Global) GetVariable(name string) builtin.Object {
	return g.Variables[name]
}

func (g *Global) SetVariable(name string, value builtin.Object) {
	g.Variables[name] = value
}

func (g *Global) RegisterFun(name string, fun Fun) {
	g.Functions[name] = fun
}

func (g *Global) ExecuteFun(name string, args []builtin.Object) builtin.Object {
	return g.Functions[name](args)
}

func (g *Global) GetFun(name string) Fun {
	if g.Functions[name] == nil {
		return nil
	}
	return g.Functions[name]
}
