package neko

import (
	ast "github.com/hedarikun/neko/ast"
	builtin "github.com/hedarikun/neko/builtin"
	inputStream "github.com/hedarikun/neko/inputStream"
	lexer "github.com/hedarikun/neko/lexer"
	parser "github.com/hedarikun/neko/parser"
)

type Evaluator struct {
	Global Global
}

func New() *Evaluator {
	return &Evaluator{
		Global: Global{
			Variables: make(map[string]builtin.Object, 0),
			Functions: make(map[string]builtin.FunObject, 0),
		},
	}
}

func (e *Evaluator) StartEvaluate(context string) {
	ins := inputStream.New(context)
	lex := lexer.New(ins)
	lex.Lexerize()
	p := parser.New(lex)
	prog := p.Parse()
	//  fmt.Printf("%+v", prog.Statements)
	for _, statement := range prog.Statements {
		evaluateStatement(statement, &e.Global)
	}

}

func evaluateStatement(statement ast.Statement, scope ScopeInterface) builtin.Object {
	switch val := statement.(type) {
	case ast.ExpressionStatment:
		return evaluateExpression(val.Value, scope)
	case ast.LetStatment:
		value := evaluateExpression(val.Value, scope)
		if value != nil {
			scope.SetVariable(val.Name.Value, value)
		}
	}

	return nil
}

func evaluateExpression(val ast.Expression, scope ScopeInterface) builtin.Object {
	switch val := val.(type) {
	case ast.OperationExpression:
		return evaluateOperationExpression(val, scope)
	case ast.NumberExpression:
		return evaluateNumber(val)
	case ast.StringExpression:
		return evaluateString(val)
	case ast.BoolExpression:
		return evaluateBool(val)
	case ast.ArrayExpression:
		return evaluateArray(val, scope)
	case ast.CallExpression:
		return evaluateCallExpression(val, scope)
	case ast.IfExpression:
		return evaluateIfExpression(val, scope)
	case ast.Identifier:
		return evaluateIdentifier(val, scope)
	case ast.FunExpression:
		return evaluateFunExpression(val, scope)
	case ast.ArrayCallExpression:
		return evaluateArrayExpreesion(val, scope)
	default:
		return nil
	}
}

func evaluateArrayExpreesion(val ast.ArrayCallExpression, scope ScopeInterface) builtin.Object {
	variable := scope.GetVariable(val.Ident.Value)
	if variable == nil {
		// error handling
	}
	fun := variable.GetMethod("indexOf")
	if fun == nil {
		// error handling
	}
	exp := evaluateExpression(val.Index, scope)
	return fun([]builtin.Object{exp})
}

func evaluateFunExpression(val ast.FunExpression, scope ScopeInterface) builtin.Object {
	body := func(args []builtin.Object) builtin.Object {
		if len(args) < len(val.Parameters) {
			// error handling
		}

		innerScope := NewScope()
		innerScope.Outer = scope

		for i, arg := range val.Parameters {
			innerScope.SetVariable(arg.Value, args[i])
		}

		block, _ := val.Body.(ast.BlockExpression)
		return evaluateBlockExpression(block, innerScope)
	}

	if val.Name.Value != "" {
		scope.RegisterFun(val.Name.Value, body)
		return nil
	} else {
		return builtin.NewFun(body)
	}
}

func evaluateIdentifier(val ast.Identifier, scope ScopeInterface) builtin.Object {
	if v := scope.GetVariable(val.Value); v != nil {
		return v
	}
	if f := scope.GetFun(val.Value); f != nil {
		return f
	}
	// error identifier doesn't exist ToDo
	return nil
}

func evaluateIfExpression(val ast.IfExpression, scope ScopeInterface) builtin.Object {
	exp := evaluateExpression(val.Condition, scope)
	boolean, ok := exp.(builtin.BoolObject)
	if !ok {
		// error
	}
	innerScope := NewScope()
	innerScope.Outer = scope

	if boolean.Value == true {
		ifBlock, _ := val.IfBlock.(ast.BlockExpression)
		return evaluateBlockExpression(ifBlock, innerScope)
	} else {
		elseBlock, _ := val.ElseBlock.(ast.BlockExpression)
		return evaluateBlockExpression(elseBlock, innerScope)
	}
}

func evaluateBlockExpression(val ast.BlockExpression, scope ScopeInterface) builtin.Object {
	if len(val.Statements) <= 0 {
		return nil
	}
	for i := 0; i < len(val.Statements)-1; i++ {
		evaluateStatement(val.Statements[i], scope)
	}
	return evaluateStatement(val.Statements[len(val.Statements)-1], scope)
}

func evaluateCallExpression(val ast.CallExpression, scope ScopeInterface) builtin.Object {
	var args []builtin.Object
	for _, arg := range val.Args {
		val := evaluateExpression(arg, scope)
		args = append(args, val)
	}

	if f := scope.GetVariable(val.Ident.Value); f != nil {
		if fun := f.GetMethod("call"); fun != nil {
			return fun(args)
		}
	}

	if f := scope.GetFun(val.Ident.Value); f != nil {
		if fun := f.GetMethod("call"); fun != nil {
			return fun(args)
		}
	}
	//ToDo error handling
	return nil
}

func evaluateOperationExpression(val ast.OperationExpression, scope ScopeInterface) builtin.Object {
	left := evaluateExpression(val.Left, scope)
	right := evaluateExpression(val.Right, scope)
	switch val.Operator.Type {
	case lexer.PLUS:
		return left.CallMethod("add", []builtin.Object{right})
	case lexer.MINUS:
		return left.CallMethod("subtract", []builtin.Object{right})
	case lexer.MULTI:
		return left.CallMethod("multiply", []builtin.Object{right})
	case lexer.DIVIDE:
		return left.CallMethod("divide", []builtin.Object{right})
	case lexer.EQUAL:
		return left.CallMethod("equal", []builtin.Object{right})
	case lexer.GREATER:
		return left.CallMethod("greater", []builtin.Object{right})
	case lexer.GREATER_OR_EQUAL:
		return left.CallMethod("greaterOrEqual", []builtin.Object{right})
	case lexer.LOWER:
		return left.CallMethod("lower", []builtin.Object{right})
	case lexer.LOWER_OR_EQUAL:
		return left.CallMethod("lowerOrEqual", []builtin.Object{right})
	case lexer.AND:
		return left.CallMethod("and", []builtin.Object{right})
	case lexer.OR:
		return left.CallMethod("or", []builtin.Object{right})
	default:
		return nil
	}

}

func evaluateNumber(val ast.NumberExpression) builtin.Object {
	return builtin.NewNumber(val.Value)
}

func evaluateString(val ast.StringExpression) builtin.Object {
	return builtin.NewString(val.Value)
}

func evaluateBool(val ast.BoolExpression) builtin.Object {
	return builtin.NewBool(val.Value)
}

func evaluateArray(val ast.ArrayExpression, scope ScopeInterface) builtin.Object {
	var objects []builtin.Object
	for _, exp := range val.Values {
		expVal := evaluateExpression(exp, scope)
		if expVal == nil {
			// error handling
		}
		objects = append(objects, expVal)
	}
	return builtin.NewArray(objects)
}
