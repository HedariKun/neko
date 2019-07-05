package neko

import (
	builtin "github.com/hedarikun/neko/builtin"
)

type ScopeInterface interface {
	GetVariable(string) builtin.Object
	SetVariable(string, builtin.Object)
	RegisterFun(string, func([]builtin.Object) builtin.Object)
	ExecuteFun(string, []builtin.Object) builtin.Object
}

type Scope struct {
	Functions map[string]func([]builtin.Object) builtin.Object
	Variables map[string]builtin.Object
	Outer     ScopeInterface
}

func (s *Scope) GetVariable(name string) builtin.Object {
	return s.Variables[name]
}

func (s *Scope) SetVariable(name string, value builtin.Object) {
	s.Variables[name] = value
}

func (s *Scope) RegisterFun(name string, fun func([]builtin.Object) builtin.Object) {
	s.Functions[name] = fun
}

func (s *Scope) ExecuteFun(name string, args []builtin.Object) builtin.Object {
	return s.Functions[name](args)
}

type Global struct {
	Functions map[string]func([]builtin.Object) builtin.Object
	Variables map[string]builtin.Object
}

func (g *Global) GetVariable(name string) builtin.Object {
	return g.Variables[name]
}

func (g *Global) SetVariable(name string, value builtin.Object) {
	g.Variables[name] = value
}

func (g *Global) RegisterFun(name string, fun func([]builtin.Object) builtin.Object) {
	g.Functions[name] = fun
}

func (g *Global) ExecuteFun(name string, args []builtin.Object) builtin.Object {
	return g.Functions[name](args)
}
