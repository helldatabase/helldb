package ast

import (
	"bytes"
	"fmt"
	"helldb/query/token"
	"strings"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type ValueExpression interface {
	Node
	expressionNode()
}

type Query struct {
	Statements []Statement
}

func (q *Query) TokenLiteral() string {
	if len(q.Statements) > 0 {
		return q.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (q *Query) String() string {
	var out bytes.Buffer

	for _, s := range q.Statements {
		out.WriteString(s.String() + "\n")
	}

	return out.String()
}

type Identifier struct {
	Token token.Token // the IDENT token
	Value string
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) String() string       { return i.Value }
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

type GetStatement struct {
	Token token.Token  // the GET token
	Keys  []Identifier // array of keys to retrieve
}

func (gs *GetStatement) statementNode()       {}
func (gs *GetStatement) TokenLiteral() string { return gs.Token.Literal }
func (gs *GetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(gs.TokenLiteral() + " ")

	if len(gs.Keys) == 1 {
		out.WriteString(gs.Keys[0].String())
	} else {
		for i := 0; i < len(gs.Keys)-1; i++ {
			out.WriteString(gs.Keys[i].String() + " & ")
		}
		out.WriteString(gs.Keys[len(gs.Keys)-1].String())
	}

	out.WriteString(";")

	return out.String()
}

type PutStatement struct {
	Token token.Token     // the token.PUT token
	Key   Identifier      // the key to set to
	Value ValueExpression // the value to store
}

func (ps *PutStatement) statementNode()       {}
func (ps *PutStatement) TokenLiteral() string { return ps.Token.Literal }
func (ps *PutStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ps.TokenLiteral() + " ")
	out.WriteString(ps.Key.String() + " ")

	if ps.Value != nil {
		out.WriteString(ps.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

type DelStatement struct {
	Token token.Token  // the DEL token
	Keys  []Identifier // array of keys to delete
}

func (ds *DelStatement) statementNode()       {}
func (ds *DelStatement) TokenLiteral() string { return ds.Token.Literal }
func (ds *DelStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")

	if len(ds.Keys) == 1 {
		out.WriteString(ds.Keys[0].String())
	} else {
		for i := 0; i < len(ds.Keys)-1; i++ {
			out.WriteString(ds.Keys[i].String() + " & ")
		}
		out.WriteString(ds.Keys[len(ds.Keys)-1].String())
	}

	out.WriteString(";")

	return out.String()
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }
func (il *IntegerLiteral) String() string       { return il.Token.Literal }

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode()      {}
func (sl *StringLiteral) TokenLiteral() string { return sl.Token.Literal }
func (sl *StringLiteral) String() string {
	return fmt.Sprintf(`"%s"`, sl.Token.Literal)
}

type BooleanLiteral struct {
	Token token.Token
	Value bool
}

func (bl *BooleanLiteral) expressionNode()      {}
func (bl *BooleanLiteral) TokenLiteral() string { return bl.Token.Literal }
func (bl *BooleanLiteral) String() string {
	if bl.Value {
		return "true"
	} else {
		return "false"
	}
}

type CollectionLiteral struct {
	Token    token.Token
	Elements []ValueExpression
}

func (cl *CollectionLiteral) expressionNode()      {}
func (cl *CollectionLiteral) TokenLiteral() string { return cl.Token.Literal }
func (cl *CollectionLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, el := range cl.Elements {
		elements = append(elements, el.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
