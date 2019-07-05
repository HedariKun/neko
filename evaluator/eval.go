package neko

import (
	ast "github.com/hedarikun/neko/ast"
	inputStream "github.com/hedarikun/neko/inputStream"
	lexer "github.com/hedarikun/neko/lexer"
	parser "github.com/hedarikun/neko/parser"
	builtin "github.com/hedarikun/neko/builtin"
	"fmt"
)

type Evaluator struct {
	Global Global
}

func New() *Evaluator {
	return &Evaluator{
		Global: Global{
			Variables: make(map[string]builtin.Object, 0),
			Functions: make(map[string]func([]builtin.Object) builtin.Object, 0),
		},
	}
}

func (e *Evaluator) StartEvaluate(context string) {
	ins := inputStream.New(context)
	lex := lexer.New(ins)
	lex.Lexerize()
	p := parser.New(lex)
	prog := p.Parse()
	fmt.Printf("%+v", prog.Statements)
	// for _, statement := range prog.Statements {
	// 	switch val := statement.(type) {
	// 	case ast.ExpressionStatment:
	// 		evaluateExpression(val.Value, &e.Global)
	// 	case ast.LetStatment:
	// 		value := evaluateExpression(val.Value, &e.Global)
	// 		if value != nil {
	// 			e.Global.SetVariable(val.Name.Value, value)
	// 		}
	// 	}
	// }
}

func evaluateExpression(val ast.Expression, scope ScopeInterface) builtin.Object {
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

func evaluateCallExpression(val ast.CallExpression, scope ScopeInterface) builtin.Object {
	var args []builtin.Object
	for _, arg := range val.Args {
		args = append(args, evaluateExpression(arg, scope))
	}
	return scope.ExecuteFun(val.Ident.Value, args)
}

func evaluateOperationExpression(val ast.OperationExpression, scope ScopeInterface) builtin.Object {
	left := evaluateExpression(val.Left, scope)
	right := evaluateExpression(val.Right, scope)
	switch val.Operator.Type {
	case lexer.PLUS:
		return left.CallMethod("add", []builtin.Object{right})
	case lexer.MULTI:
		return left.CallMethod("subtract", []builtin.Object{right})
	case lexer.MINUS:
		return left.CallMethod("multiply", []builtin.Object{right})
	case lexer.DIVIDE:
		return left.CallMethod("divide", []builtin.Object{right})
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
