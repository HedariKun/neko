package neko

import (
	inputStream "github.com/hedarikun/neko/inputStream"
)

type Lexer struct {
	inputStream *inputStream.InputStream
	token       []Token
	Position    int
}

func (t *Lexer) Lexerize() {
	for !t.inputStream.EOF() {
		switch char := t.inputStream.Peek(); char {
		case '!':
			t.AddToken(NewToken(BANG, string(t.inputStream.Next())))
		case '-':
			t.AddToken(NewToken(MINUS, string(t.inputStream.Next())))
		case '+':
			t.AddToken(NewToken(PLUS, string(t.inputStream.Next())))
		case '*':
			t.AddToken(NewToken(MULTI, string(t.inputStream.Next())))
		case '/':
			t.AddToken(NewToken(DIVIDE, string(t.inputStream.Next())))
		case '.':
			t.AddToken(NewToken(DOT, string(t.inputStream.Next())))
		case ',':
			t.AddToken(NewToken(COMMA, string(t.inputStream.Next())))
		case '|':
			char := t.inputStream.Next()
			if t.inputStream.Peek() == '|' {
				t.AddToken(NewToken(OR, string(string(char)+string(t.inputStream.Next()))))
			} else {
				// error
			}
		case '&':
			char := t.inputStream.Next()
			if t.inputStream.Peek() == '&' {
				t.AddToken(NewToken(AND, string(string(char)+string(t.inputStream.Next()))))
			} else {
				// error
			}
		case '[':
			t.AddToken(NewToken(OB, string(t.inputStream.Next())))
		case ']':
			t.AddToken(NewToken(CB, string(t.inputStream.Next())))
		case '{':
			t.AddToken(NewToken(OCB, string(t.inputStream.Next())))
		case '}':
			t.AddToken(NewToken(CCB, string(t.inputStream.Next())))
		case '(':
			t.AddToken(NewToken(OP, string(t.inputStream.Next())))
		case ')':
			t.AddToken(NewToken(CP, string(t.inputStream.Next())))
		case '>':
			char := t.inputStream.Next()
			if t.inputStream.Peek() == '=' {
				t.AddToken(NewToken(GREATER_OR_EQUAL, string(char)+string(t.inputStream.Next())))
			} else {
				t.AddToken(NewToken(GREATER, string(char)))
			}
		case '<':
			char := t.inputStream.Next()
			if t.inputStream.Peek() == '=' {
				t.AddToken(NewToken(LOWER_OR_EQUAL, string(char)+string(t.inputStream.Next())))
			} else {
				t.AddToken(NewToken(LOWER, string(char)))
			}
		case '=':
			char := t.inputStream.Next()
			if t.inputStream.Peek() == '=' {
				t.AddToken(NewToken(EQUAL, string(char)+string(t.inputStream.Next())))
			} else {
				t.AddToken(NewToken(ASSIGN, string(char)))
			}
		case ';':
			t.AddToken(NewToken(SEMI, string(t.inputStream.Next())))
		case ' ':
			t.inputStream.Next()
		case '\n':
			t.inputStream.Next()
		case '\r':
			t.inputStream.Next()
		default:
			if isNumber(char) {
				token := lexerizeNumber(t)
				t.AddToken(token)
			} else if char == '"' || char == '\'' || char == '`' {
				token := lexerizeString(t)
				t.AddToken(token)
			} else if isLetter(char) || char == '_' {
				token := lexerizeKeyword(t)
				t.AddToken(token)
			} else {
				// To Do: better error handling.
				t.inputStream.Next()
			}
		}
	}
	t.AddToken(Token{Type: EOF, Value: ""})
}

func (l *Lexer) AddToken(token Token) {
	l.token = append(l.token, token)
}

func lexerizeString(tokenizer *Lexer) Token {
	origin := tokenizer.inputStream.Next()
	val := string(tokenizer.inputStream.Next())
	for tokenizer.inputStream.Peek() != origin {
		val += string(tokenizer.inputStream.Next())
	}
	tokenizer.inputStream.Next()
	return NewToken(STRING, val)
}

func lexerizeNumber(tokenizer *Lexer) Token {
	val := string(tokenizer.inputStream.Next())
	isFloat := false
	for isNumber(tokenizer.inputStream.Peek()) || !isFloat && tokenizer.inputStream.Peek() == '.' {
		if tokenizer.inputStream.Peek() == '.' {
			isFloat = true
		}
		val += string(tokenizer.inputStream.Next())
	}
	return NewToken(NUMBER, val)
}

func lexerizeKeyword(tokenizer *Lexer) Token {
	val := string(tokenizer.inputStream.Next())
	for isNumber(tokenizer.inputStream.Peek()) || isLetter(tokenizer.inputStream.Peek()) || tokenizer.inputStream.Peek() == '_' {
		val += string(tokenizer.inputStream.Next())
	}
	t := IDENT
	switch val {
	case LET:
		t = LET
	case FUN:
		t = FUN
	case IF:
		t = IF
	case ELSE:
		t = ELSE
	case FALSE:
		t = BOOL
	case TRUE:
		t = BOOL
	case MUT:
		t = MUT
	case STRUCT:
		t = STRUCT
	}
	return NewToken(t, val)
}

func (l *Lexer) Next() Token {
	token := l.token[l.Position]
	l.Position++
	return token
}

func (l *Lexer) Peek() Token {
	token := l.token[l.Position]
	return token
}

func (l *Lexer) PeekFurther() Token {
	token := l.token[l.Position+1]
	return token
}

func (l *Lexer) EOF() bool {
	return l.Position >= len(l.token) || l.Peek().Type == EOF
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func isLetter(char rune) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z'
}

func New(is *inputStream.InputStream) *Lexer {
	return &Lexer{
		inputStream: is,
	}
}
