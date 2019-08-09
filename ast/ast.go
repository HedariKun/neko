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

type LetStatement struct {
	Token lexer.Token
	Mut   bool
	Name  Identifier
	Value Expression
}

func (ls LetStatement) TokenLiteral() string {
	return ls.Token.Type
}

func (ls LetStatement) statementNode() {}

type StructStatement struct {
	Token lexer.Token
	Name  Identifier
	Props []Identifier
}

func (ss StructStatement) TokenLiteral() string {
	return ss.Token.Type
}

func (ss StructStatement) statementNode() {}

type AssignmentExpression struct {
	Token lexer.Token
	Ident Identifier
	Value Expression
}

func (ae AssignmentExpression) TokenLiteral() string {
	return ae.Token.Value
}

func (ae AssignmentExpression) expressionNode() {}

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

type StringExpression struct {
	Token lexer.Token
	Value string
}

func (se StringExpression) TokenLiteral() string {
	return se.Token.Type
}

func (se StringExpression) expressionNode() {}

type BoolExpression struct {
	Token lexer.Token
	Value bool
}

func (be BoolExpression) TokenLiteral() string {
	return be.Token.Type
}

func (be BoolExpression) expressionNode() {}

type ArrayExpression struct {
	Token  lexer.Token
	Values []Expression
}

func (ae ArrayExpression) TokenLiteral() string {
	return ae.Token.Value
}

func (ae ArrayExpression) expressionNode() {}

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

type CallExpression struct {
	Token  lexer.Token
	Object Expression
	Args   []Expression
}

func (ce CallExpression) TokenLiteral() string {
	return ce.Token.Value
}

func (ce CallExpression) expressionNode() {}

type ArrayCallExpression struct {
	Token  lexer.Token
	Object Expression
	Index  Expression
}

func (ae ArrayCallExpression) TokenLiteral() string {
	return ae.Token.Value
}

func (ae ArrayCallExpression) expressionNode() {}

type FieldCallExpression struct {
	Token  lexer.Token
	Object Expression
	Child  Identifier
}

func (fe FieldCallExpression) TokenLiteral() string {
	return fe.Token.Value
}

func (fe FieldCallExpression) expressionNode() {}
