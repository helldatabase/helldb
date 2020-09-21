package lexer

import "helldb/query/token"

type Lexer struct {
	ch           byte
	input        string
	position     int
	readPosition int
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {

	var tokenType token.Type

	l.skipWhitespace()

	switch l.ch {
	case '[':
		tokenType = token.LBRACK
	case ']':
		tokenType = token.RBRACK
	case ';':
		tokenType = token.SEMICOLON
	case ',':
		tokenType = token.COMMA
	case '&':
		tokenType = token.AND
	case '"':
		tok := token.Token{Type: token.STRING}
		tok.Literal = l.readString()
		return tok
	case 0:
		l.readChar()
		return token.Token{Type: token.EOF, Literal: ""}
	default:
		return parseMultiCharString(l)
	}

	tok := token.New(tokenType, l.ch)
	l.readChar()
	return tok

}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func parseMultiCharString(l *Lexer) token.Token {
	var tok token.Token
	if isLetter(l.ch) {
		tok.Literal = l.readIdentifier()
		tok.Type = token.LookupIdent(tok.Literal)
		return tok
	} else if isDigit(l.ch) {
		tok.Literal = l.readNumber()
		tok.Type = token.INT
		return tok
	} else {
		ch := l.ch
		l.readChar()
		return token.New(token.ILLEGAL, ch)
	}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	defer l.readChar()
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
