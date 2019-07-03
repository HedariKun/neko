package neko

import (
	lexer "github.com/hedarikun/neko/lexer"
)

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type LetStatment struct {
	Token lexer.Token
	Name  Identifier
	Value Expression
}

func (ls LetStatment) TokenLiteral() string {
	return ls.Token.Type
}

func (ls LetStatment) statementNode() {}

type BlockExpression struct {
	Statements []Statement
}

func (bs BlockExpression) TokenLiteral() string {
	if len(bs.Statements) > 0 {
		return bs.Statements[0].TokenLiteral()
	}
	return ""
}

func (bs BlockExpression) expressionNode() {}

type ExpressionStatment struct {
	Value Expression
}

func (es ExpressionStatment) TokenLiteral() string {
	return es.Value.TokenLiteral()
}

func (es ExpressionStatment) statementNode() {}

type PrefixExpression struct {
	Prefix lexer.Token
	Value  Expression
}

func (pe PrefixExpression) TokenLiteral() string {
	return pe.Prefix.Type
}

func (pe PrefixExpression) expressionNode() {}

type NumberExpression struct {
	Token lexer.Token
	Value float64
}

func (ne NumberExpression) TokenLiteral() string {
	return ne.Token.Type
}

func (ne NumberExpression) expressionNode() {}

type OperationExpression struct {
	Operator lexer.Token
	Left     Expression
	Right    Expression
}

func (oe OperationExpression) TokenLiteral() string {
	return oe.Operator.Type
}

func (oe OperationExpression) expressionNode() {}

type IfExpression struct {
	Condition Expression
	IfBlock   Expression
	ElseBlock Expression
}

func (ie IfExpression) TokenLiteral() string {
	return ie.Condition.TokenLiteral()
}

func (ie IfExpression) expressionNode() {}

type Identifier struct {
	Token lexer.Token
	Value string
}

func (ie Identifier) TokenLiteral() string {
	return ie.Token.Type
}

func (ie Identifier) expressionNode() {}

type FunExpression struct {
	Token      lexer.Token
	Name       Identifier
	Parameters []Identifier
	Body       Expression
}

func (fe FunExpression) TokenLiteral() string {
	return fe.Token.Type
}

func (fe FunExpression) expressionNode() {}