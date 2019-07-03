package neko

import (
	"strconv"

	ast "github.com/hedarikun/neko/ast"
	lexer "github.com/hedarikun/neko/lexer"
)

type Parser struct {
	l *lexer.Lexer
}

func (p *Parser) Parse() *ast.Program {
	program := ast.Program{}

	for !p.l.EOF() {
		statement := parseStatement(p.l)
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
	}

	return &program
}

func parseStatement(l *lexer.Lexer) ast.Statement {
	switch l.Peek().Type {
	case lexer.LET:
		return parseLetStatement(l)
	default:
		return parseExpressionStatment(l)
	}
}

func parseExpressionStatment(l *lexer.Lexer) ast.Statement {
	return ast.ExpressionStatment{
		Value: parseExpression(l),
	}
}

func parseLetStatement(l *lexer.Lexer) ast.Statement {
	literal := l.Next()

	if l.Peek().Type != lexer.IDENT {
		// error handling
	}
	name := l.Next()

	if l.Peek().Type != lexer.ASSIGN {
		// error handling
	}
	l.Next()

	val := parseExpression(l)

	return ast.LetStatment{
		Token: literal,
		Name: ast.Identifier{
			Token: name,
			Value: name.Value,
		},
		Value: val,
	}
}

func parseBlockExpression(l *lexer.Lexer) ast.Expression {
	l.Next()
	blockExpression := ast.BlockExpression{}

	for !l.EOF() && l.Peek().Type != lexer.CCB {
		statement := parseStatement(l)
		blockExpression.Statements = append(blockExpression.Statements, statement)
	}

	if !l.EOF() {
		// error
	}
	l.Next()

	return blockExpression
}

func parseExpression(l *lexer.Lexer) ast.Expression {
	var leftExpression ast.Expression

	switch l.Peek().Type {
	case lexer.NUMBER:
		leftExpression = parseNumberExpression(l)
	case lexer.IF:
		leftExpression = parseIfExpression(l)
	case lexer.OCB:
		leftExpression = parseBlockExpression(l)
	case lexer.FUN:
		leftExpression = parseFunExpression(l)
	default:
		// undefined: error handle
	}

	return leftExpression
}

func parseNumberExpression(l *lexer.Lexer) ast.Expression {
	token := l.Next()

	value, err := strconv.ParseFloat(token.Value, 64)
	if err != nil {
		// error handle
	}
	var numberExpression ast.Expression
	numberExpression = ast.NumberExpression{
		Token: token,
		Value: value,
	}

	if !l.EOF() && isOperation(l.Peek()) {
		numberExpression = parseOperationExpression(l, numberExpression)
	}
	return numberExpression
}

func parseOperationExpression(l *lexer.Lexer, leftExpression ast.Expression) ast.Expression {
	operation := l.Next()
	rightSide := parseExpression(l)
	if !l.EOF() && isOperation(l.Peek()) && (operation.Type == lexer.PLUS || operation.Type == lexer.MINUS) {
		if l.Peek().Type == lexer.MULTI || l.Peek().Type == lexer.MULTI {
			rightSide = parseOperationExpression(l, rightSide)
		} else {
			curOp := ast.OperationExpression{Operator: operation, Left: leftExpression, Right: rightSide}
			nextOp := parseOperationExpression(l, curOp)
			return nextOp
		}
	}
	return ast.OperationExpression{
		Operator: operation,
		Left:     leftExpression,
		Right:    rightSide,
	}
}

func parsePrefixExpression(l *lexer.Lexer) ast.Expression {
	prefix := l.Next()
	val := parseExpression(l)

	return ast.PrefixExpression{
		Prefix: prefix,
		Value:  val,
	}
}

func parseIfExpression(l *lexer.Lexer) ast.Expression {
	ifExpression := ast.IfExpression{}
	l.Next()
	if l.EOF() {
		// error
	}

	ifExpression.Condition = parseExpression(l)

	if l.EOF() || l.Peek().Type != lexer.OCB {
		// error
	}

	ifExpression.IfBlock = parseExpression(l)

	if !l.EOF() && l.Peek().Type == lexer.ELSE {
		l.Next()
		if l.EOF() || l.Peek().Type != lexer.OCB {
			// error
		}
		ifExpression.ElseBlock = parseExpression(l)
	}

	return ifExpression
}

func parseFunExpression(l *lexer.Lexer) ast.Expression {
	funExp := ast.FunExpression{
		Token: l.Next(),
	}

	if l.EOF() || l.Peek().Type != lexer.IDENT {
		// error
	}

	identToken := l.Next()
	identifier := ast.Identifier{
		Token: identToken,
		Value: identToken.Value,
	}

	funExp.Name = identifier

	if l.EOF() || l.Peek().Type != lexer.OP {
		// error
	}
	l.Next()

	funExp.Parameters = parseParameters(l)

	if l.EOF() || l.Peek().Type != lexer.OCB {
		// error
	}

	funExp.Body = parseBlockExpression(l)

	return funExp
}

func parseParameters(l *lexer.Lexer) []ast.Identifier {
	var parameters []ast.Identifier
	for !l.EOF() && l.Peek().Type != lexer.CP {
		if l.Peek().Type == lexer.IDENT {
			t := l.Next()
			parameters = append(parameters, ast.Identifier{
				Token: t,
				Value: t.Value,
			})
			if l.Peek().Type == lexer.IDENT {
				// error
			}
		} else if l.Peek().Type == lexer.Comma {
			l.Next()
		} else {
			// error
		}
	}
	if l.EOF() {
		// error
	}
	l.Next()
	return parameters
}

func isPrefix(t lexer.Token) bool {
	return t.Type == lexer.BANG || t.Type == lexer.MINUS
}

func isOperation(t lexer.Token) bool {
	return t.Type == lexer.PLUS || t.Type == lexer.MULTI || t.Type == lexer.MINUS || t.Type == lexer.DIVIDE
}

func New(l *lexer.Lexer) *Parser {
	return &Parser{
		l: l,
	}
}
