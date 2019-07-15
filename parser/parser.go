package neko

import (
	"strconv"

	ast "github.com/hedarikun/neko/ast"
	lexer "github.com/hedarikun/neko/lexer"
)

type Parser struct {
	l *lexer.Lexer
}

var operLevels map[string]int

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
		return parseExpressionStatement(l)
	}
}

func parseExpressionStatement(l *lexer.Lexer) ast.Statement {
	return ast.ExpressionStatment{
		Value: parseExpression(l, 0),
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

	val := parseExpression(l, 0)

	return ast.LetStatment{
		Token: literal,
		Name: ast.Identifier{
			Token: name,
			Value: name.Value,
		},
		Value: val,
	}
}

func parseExpression(l *lexer.Lexer, curOperLevel int) ast.Expression {
	var leftExpression ast.Expression

	switch l.Peek().Type {
	case lexer.NUMBER:
		leftExpression = parseNumberExpression(l)
	case lexer.STRING:
		leftExpression = parseStringExpression(l)
	case lexer.BOOL:
		leftExpression = parseBoolExpression(l)
	case lexer.OB:
		leftExpression = parseArrayExpression(l)
	case lexer.IF:
		leftExpression = parseIfExpression(l)
	case lexer.OCB:
		leftExpression = parseBlockExpression(l)
	case lexer.FUN:
		leftExpression = parseFunExpression(l)
	case lexer.IDENT:
		leftExpression = parseIdentifier(l)
	default:
		// undefined: error handle
	}

	for !l.EOF() && isOperation(l.Peek()) && (operLevels[l.Peek().Type] < curOperLevel || curOperLevel == 0) {
		leftExpression = parseOperationExpression(l, leftExpression)
	}

	return leftExpression
}

func parseIdentifier(l *lexer.Lexer) ast.Expression {
	t := l.Next()
	ident := ast.Identifier{Token: t, Value: t.Value}
	switch l.Peek().Type {
	case lexer.OP:
		return parseFunCall(l, ident)
	case lexer.OB:
		return parseArrayCall(l, ident)
	}
	return ident
}

func parseFunCall(l *lexer.Lexer, ident ast.Identifier) ast.Expression {
	l.Next()
	var args []ast.Expression
	if l.Peek().Type != lexer.CP {
		args = parseArgs(l)
	}
	l.Next()
	return ast.CallExpression{
		Token: ident.Token,
		Ident: ident,
		Args:  args,
	}
}

func parseArrayCall(l *lexer.Lexer, ident ast.Identifier) ast.Expression {
	l.Next()
	var index ast.Expression
	if l.Peek().Type != lexer.CB {
		index = parseExpression(l, 0)
	}

	if index == nil {
		// error
	}
	l.Next()

	return ast.ArrayCallExpression{
		Token: ident.Token,
		Ident: ident,
		Index: index,
	}
}

func parseArgs(l *lexer.Lexer) []ast.Expression {
	var args []ast.Expression
	for !l.EOF() && l.Peek().Type != lexer.CP {
		expression := parseExpression(l, 0)
		if l.EOF() || l.Peek().Type != lexer.COMMA {
			// error
		}
		if l.Peek().Type == lexer.COMMA {
			l.Next()
		}
		args = append(args, expression)
	}
	if l.EOF() {
		// error
	}
	return args
}

func parseBlockExpression(l *lexer.Lexer) ast.Expression {
	l.Next()
	blockExpression := ast.BlockExpression{}

	for !l.EOF() && l.Peek().Type != lexer.CCB {
		statement := parseStatement(l)
		blockExpression.Statements = append(blockExpression.Statements, statement)
	}
	if l.EOF() {
		// error
	}
	l.Next()

	return blockExpression
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

	return numberExpression
}

func parseStringExpression(l *lexer.Lexer) ast.Expression {
	token := l.Next()
	return ast.StringExpression{
		Token: token,
		Value: token.Value,
	}
}

func parseBoolExpression(l *lexer.Lexer) ast.Expression {
	token := l.Next()
	value := false
	if token.Value == "true" {
		value = true
	}
	return ast.BoolExpression{
		Token: token,
		Value: value,
	}
}

func parseArrayExpression(l *lexer.Lexer) ast.Expression {
	token := l.Next()
	var values []ast.Expression
	for l.Peek().Type != lexer.CB {
		exp := parseExpression(l, 0)
		values = append(values, exp)
		if l.Peek().Type != lexer.CB && l.Peek().Type != lexer.COMMA {
			// error handling
		}
		if l.Peek().Type == lexer.COMMA {
			l.Next()
		}
	}
	l.Next()
	return ast.ArrayExpression{
		Token:  token,
		Values: values,
	}
}

func parseOperationExpression(l *lexer.Lexer, leftExpression ast.Expression) ast.Expression {
	operation := l.Next()
	rightSide := parseExpression(l, operLevels[operation.Type])

	if !l.EOF() && isOperation(l.Peek()) {
		if operLevels[operation.Type] > operLevels[l.Peek().Type] {
			rightSide = parseOperationExpression(l, rightSide)
		}
	}
	var curOp ast.Expression
	curOp = ast.OperationExpression{Operator: operation, Left: leftExpression, Right: rightSide}
	return curOp
}

func parsePrefixExpression(l *lexer.Lexer) ast.Expression {
	prefix := l.Next()
	val := parseExpression(l, 0)

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

	ifExpression.Condition = parseExpression(l, 0)
	if l.EOF() || l.Peek().Type != lexer.OCB {
		// error
	}

	ifExpression.IfBlock = parseExpression(l, 0)

	if !l.EOF() && l.Peek().Type == lexer.ELSE {
		l.Next()
		if l.EOF() || l.Peek().Type != lexer.OCB {
			// error
		}
		ifExpression.ElseBlock = parseExpression(l, 0)
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
		} else if l.Peek().Type == lexer.COMMA {
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
	return t.Type == lexer.PLUS || t.Type == lexer.MULTI || t.Type == lexer.MINUS || t.Type == lexer.DIVIDE || t.Type == lexer.GREATER || t.Type == lexer.GREATER_OR_EQUAL || t.Type == lexer.LOWER || t.Type == lexer.LOWER_OR_EQUAL || t.Type == lexer.EQUAL || t.Type == lexer.AND || t.Type == lexer.OR
}

func New(l *lexer.Lexer) *Parser {
	operLevels = make(map[string]int, 0)
	operLevels[lexer.MULTI] = 1
	operLevels[lexer.DIVIDE] = 1
	operLevels[lexer.PLUS] = 2
	operLevels[lexer.MINUS] = 2
	operLevels[lexer.EQUAL] = 3
	operLevels[lexer.GREATER] = 3
	operLevels[lexer.LOWER] = 3
	operLevels[lexer.GREATER_OR_EQUAL] = 3
	operLevels[lexer.LOWER_OR_EQUAL] = 3
	operLevels[lexer.AND] = 4
	operLevels[lexer.OR] = 5

	return &Parser{
		l: l,
	}
}
