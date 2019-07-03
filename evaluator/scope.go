package neko

type ScopeInterface interface {
	GetVariable(string) Object
	SetVariable(string, Object)
	RegisterFun(string, func([]Object) Object)
	ExecuteFun(string, []Object) Object
}

type Scope struct {
	Functions map[string]func([]Object) Object
	Variables map[string]Object
	Outer     ScopeInterface
}

func (s *Scope) GetVariable(name string) Object {
	return s.Variables[name]
}

func (s *Scope) SetVariable(name string, value Object) {
	s.Variables[name] = value
}

func (s *Scope) RegisterFun(name string, fun func([]Object) Object) {
	s.Functions[name] = fun
}

func (s *Scope) ExecuteFun(name string, args []Object) Object {
	return s.Functions[name](args)
}

type Global struct {
	Functions map[string]func([]Object) Object
	Variables map[string]Object
}

func (g *Global) GetVariable(name string) Object {
	return g.Variables[name]
}

func (g *Global) SetVariable(name string, value Object) {
	g.Variables[name] = value
}

func (g *Global) RegisterFun(name string, fun func([]Object) Object) {
	g.Functions[name] = fun
}

func (g *Global) ExecuteFun(name string, args []Object) Object {
	return g.Functions[name](args)
}
