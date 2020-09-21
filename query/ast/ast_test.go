package ast

import (
	"testing"

	. "helldb/query/token"
)

func TestString(t *testing.T) {
	query := &Query{
		Statements: []Statement{
			&PutStatement{
				Token: Token{Type: PUT, Literal: "PUT"},
				Key: Identifier{
					Token: Token{Type: IDENT, Literal: "age"},
					Value: "age",
				},
				Value: &IntegerLiteral{Token: Token{Type: INT, Literal: "69"}, Value: 69},
			},
			&PutStatement{
				Token: Token{Type: PUT, Literal: "PUT"},
				Key: Identifier{
					Token: Token{Type: IDENT, Literal: "posts"},
					Value: "posts",
				},
				Value: &CollectionLiteral{
					Token: Token{Type: LBRACK, Literal: "["},
					Elements: []ValueExpression{
						&IntegerLiteral{Token: Token{Type: INT, Literal: "420"}, Value: 420},
						&CollectionLiteral{
							Token: Token{Type: LBRACK, Literal: "["},
							Elements: []ValueExpression{
								&StringLiteral{Token: Token{Type: STRING, Literal: "hello-world"}, Value: "hello-world"},
							},
						},
					},
				},
			},
			&GetStatement{
				Token: Token{Type: GET, Literal: "GET"},
				Keys: []Identifier{
					{
						Token: Token{Type: IDENT, Literal: "name"},
						Value: "name",
					},
					{
						Token: Token{Type: IDENT, Literal: "age"},
						Value: "age",
					},
				},
			},
			&GetStatement{
				Token: Token{Type: GET, Literal: "GET"},
				Keys: []Identifier{
					{
						Token: Token{Type: IDENT, Literal: "sex"},
						Value: "sex",
					},
				},
			},
		},
	}
	if query.String() != `PUT age 69;
PUT posts [420, ["hello-world"]];
GET name & age;
GET sex;`+"\n" {
		t.Errorf("query is not as expected")
	}
}
