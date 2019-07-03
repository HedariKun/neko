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
	BANG             = "bang"
	POINT            = "point"
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
	STRING           = "string"
	NUMBER           = "number"
	BOOL             = "bool"
	FALSE            = "false"
	TRUE             = "true"
	IDENT            = "ident"
	LET              = "let"
	FUN              = "fun"
	IF               = "if"
	ELSE             = "else"
	OCB              = "{"
	CCB              = "}"
	OP               = "("
	CP               = ")"
	Comma            = ","
)
