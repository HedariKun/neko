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

	for _, statement := range prog.Statements {
		switch val := statement.(type) {
		case ast.ExpressionStatment:
			evaluateExpression(val.Value)
		case ast.LetStatment:
			value := evaluateExpression(val.Value)
			if value != nil {
				e.Global.SetVariable(val.Name.Value, value)
			}
		}
	}
}

func evaluateExpression(val ast.Expression) Object {
	switch val := val.(type) {
	case ast.OperationExpression:
		return evaluateOperationExpression(val)
	case ast.NumberExpression:
		return evaluateNumber(val)

	default:
		return nil
	}
}

// Will change when i add operator overloading
func evaluateOperationExpression(val ast.OperationExpression) Object {
	left, _ := evaluateExpression(val.Left).(NumberObject)
	right, _ := evaluateExpression(val.Right).(NumberObject)

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
