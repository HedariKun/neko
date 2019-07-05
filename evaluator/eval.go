package neko

import (
	"fmt"

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
			Functions: make(map[string]Fun, 0),
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
	// 	evaluateStatement(statement, &e.Global)
	// }

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
	case ast.CallExpression:
		return evaluateCallExpression(val, scope)
	case ast.IfExpression:
		return evaluateIfExpression(val, scope)
	default:
		return nil
	}
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
		args = append(args, evaluateExpression(arg, scope))
	}
	return scope.GetFun(val.Ident.Value)(args)
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

func evaluateBool(val ast.BoolExpression) builtin.Object {
	return builtin.NewBool(val.Value)
}