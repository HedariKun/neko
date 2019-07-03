package neko

const (
	EOF = '•'
)

type InputStream struct {
	source   string
	Position int
	Line     int
	Column   int
}

func (is *InputStream) Next() rune {
	char := is.source[is.Position]
	is.Position++
	if char == '\n' {
		is.Line++
		is.Column = 0
	} else {
		is.Column++
	}
	return rune(char)
}

func (is *InputStream) Peek() rune {
	if (is.Position >= len(is.source)) {
		return '•'
	}
	return rune(is.source[is.Position])
}

func (is *InputStream) EOF() bool {
	return is.Peek() == '•'
}

func (is *InputStream) Error(msg string) {
	panic(msg + "(" + string(is.Line) + " : " + string(is.Column) + ")")
}

func New(source string) *InputStream {
	return &InputStream{
		source: source,
	}
}
