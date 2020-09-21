package token

const (
	GET = "GET"
	PUT = "PUT"
	DEL = "DEL"

	INT    = "INT"
	BOOL   = "BOOL"
	IDENT  = "IDENT"
	STRING = "STRING"

	AND       = "&"
	LBRACK    = "["
	COMMA     = ","
	RBRACK    = "]"
	SEMICOLON = ";"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

type Type string

type Token struct {
	Type    Type
	Literal string
}

var keywords = map[string]Type{
	"GET":   GET,
	"PUT":   PUT,
	"DEL":   DEL,
	"true":  BOOL,
	"false": BOOL,
}

func New(tt Type, ch byte) Token {
	return Token{Type: tt, Literal: string(ch)}
}

func LookupIdent(ident string) Type {
	if t, ok := keywords[ident]; ok {
		return t
	}
	return IDENT
}
