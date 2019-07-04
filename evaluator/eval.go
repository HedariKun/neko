package neko

import (
	ast "github.com/hedarikun/neko/ast"
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
			Variables: make(map[string]Object, 0),
			Functions: make(map[string]func([]Object) Object, 0),
		},
	}
}

func (e *Evaluator) StartEvaluate(context string) {
	ins := inputStream.New(context)
	lex := lexer.New(ins)
	lex.Lexerize()
	p := parser.New(lex)
	prog := p.Parse()
	// fmt.Printf("%+v", prog.Statements)
	for _, statement := range prog.Statements {
		switch val := statement.(type) {
		case ast.ExpressionStatment:
			evaluateExpression(val.Value, &e.Global)
		case ast.LetStatment:
			value := evaluateExpression(val.Value, &e.Global)
			if value != nil {
				e.Global.SetVariable(val.Name.Value, value)
			}
		}
	}
}

func evaluateExpression(val ast.Expression, scope ScopeInterface) Object {
	switch val := val.(type) {
	case ast.OperationExpression:
		return evaluateOperationExpression(val, scope)
	case ast.NumberExpression:
		return evaluateNumber(val)
	case ast.StringExpression:
		return evaluateString(val)
	case ast.CallExpression:
		return evaluateCallExpression(val, scope)
	default:
		return nil
	}
}

func evaluateCallExpression(val ast.CallExpression, scope ScopeInterface) Object {
	var args []Object
	for _, arg := range val.Args {
		args = append(args, evaluateExpression(arg, scope))
	}
	return scope.ExecuteFun(val.Ident.Value, args)
}

// Will change when i add operator overloading
func evaluateOperationExpression(val ast.OperationExpression, scope ScopeInterface) Object {
	left, _ := evaluateExpression(val.Left, scope).(NumberObject)
	right, _ := evaluateExpression(val.Right, scope).(NumberObject)

	switch val.Operator.Type {
	case lexer.PLUS:
		return NewNumber(left.Value + right.Value)
	case lexer.MULTI:
		return NewNumber(left.Value * right.Value)
	case lexer.MINUS:
		return NewNumber(left.Value - right.Value)
	case lexer.DIVIDE:
		return NewNumber(left.Value / right.Value)
	default:
		return nil
	}

}

func evaluateNumber(val ast.NumberExpression) Object {
	return NewNumber(val.Value)
}

func evaluateString(val ast.StringExpression) Object {
	return NewString(val.Value)
}
