package neko

type Token struct {
	Type  string
	Value string
}

func NewToken(t, v string) Token {
	return Token{
		Type:  t,
		Value: v,
	}
}

const (
	EOF = "EOF"

	BANG             = "bang"
	SEMI             = ";"
	PLUS             = "+"
	MINUS            = "-"
	MULTI            = "*"
	DIVIDE           = "/"
	EQUAL            = "=="
	GREATER          = ">"
	GREATER_OR_EQUAL = ">="
	LOWER            = "<"
	LOWER_OR_EQUAL   = "<="
	ASSIGN           = "="
	OR               = "||"
	AND              = "&&"
	OB               = "["
	CB               = "]"
	OCB              = "{"
	CCB              = "}"
	OP               = "("
	CP               = ")"
	COMMA            = ","
	DOT              = "."

	STRING = "string"
	NUMBER = "number"
	BOOL   = "bool"
	FALSE  = "false"
	TRUE   = "true"

	IDENT = "ident"
	LET   = "let"
	FUN   = "fun"
	IF    = "if"
	ELSE  = "else"
	MUT   = "mut"
)
