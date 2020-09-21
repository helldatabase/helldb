package lexer

import (
	"testing"

	"helldb/query/token"
)

type expectedTypeLiteral struct {
	expectedType    token.Type
	expectedLiteral string
}

func TestNextToken(t *testing.T) {

	input1 := `[],;`
	tokens1 := []expectedTypeLiteral{
		{token.LBRACK, "["},
		{token.RBRACK, "]"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	input2 := `PUT age 15;GET sex;DEL name;`
	tokens2 := []expectedTypeLiteral{
		{token.PUT, "PUT"},
		{token.IDENT, "age"},
		{token.INT, "15"},
		{token.SEMICOLON, ";"},
		{token.GET, "GET"},
		{token.IDENT, "sex"},
		{token.SEMICOLON, ";"},
		{token.DEL, "DEL"},
		{token.IDENT, "name"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	input3 := `PUT friends [12,42,33];`
	tokens3 := []expectedTypeLiteral{
		{token.PUT, "PUT"},
		{token.IDENT, "friends"},
		{token.LBRACK, "["},
		{token.INT, "12"},
		{token.COMMA, ","},
		{token.INT, "42"},
		{token.COMMA, ","},
		{token.INT, "33"},
		{token.RBRACK, "]"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	input4 := `GET name & age;`
	tokens4 := []expectedTypeLiteral{
		{token.GET, "GET"},
		{token.IDENT, "name"},
		{token.AND, "&"},
		{token.IDENT, "age"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	input5 := `PUT male true;`
	tokens5 := []expectedTypeLiteral{
		{token.PUT, "PUT"},
		{token.IDENT, "male"},
		{token.BOOL, "true"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	input6 := `PUT name "Manan";`
	tokens6 := []expectedTypeLiteral{
		{token.PUT, "PUT"},
		{token.IDENT, "name"},
		{token.STRING, "Manan"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	tests := []struct {
		input  string
		tokens []expectedTypeLiteral
	}{
		{input1, tokens1},
		{input2, tokens2},
		{input3, tokens3},
		{input4, tokens4},
		{input5, tokens5},
		{input6, tokens6},
	}

	for i, test := range tests {
		l := New(test.input)
		for ti, tt := range test.tokens {
			tok := l.NextToken()
			if tok.Type != tt.expectedType {
				t.Errorf("tests[%d], tokens[%d] - type wrong. expected=%q, got=%q",
					i, ti, tt.expectedType, tok.Type)
			}
			if tok.Literal != tt.expectedLiteral {
				t.Errorf("tests[%d], tokens[%d] - literal wrong. expected=%q, got=%q",
					i, ti, tt.expectedType, tok.Literal)
			}
		}
	}

}
