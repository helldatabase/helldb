package parser

import (
	"strconv"
	"testing"

	"helldb/query/ast"
	"helldb/query/lexer"
)

func TestGetStatement(t *testing.T) {
	input := `GET sex; GET name & age;`

	l := lexer.New(input)
	p := New(l)

	query := p.ParseQuery()
	checkParserErrors(t, p)

	if query == nil {
		t.Fatal("ParseQuery() returned nil")
	}

	if len(query.Statements) != 2 {
		t.Fatalf("query.Statements does not contain 3 statements. got=%d",
			len(query.Statements))
	}

	tests := [][]string{
		{"sex"},
		{"name", "age"},
	}

	for i, tt := range tests {
		stmt := query.Statements[i]
		if !testGetStatement(t, stmt, tt) {
			return
		}
	}

}

func TestDelStatement(t *testing.T) {
	input := `DEL sex; DEL name & age;`

	l := lexer.New(input)
	p := New(l)

	query := p.ParseQuery()
	checkParserErrors(t, p)

	if query == nil {
		t.Fatal("ParseQuery() returned nil")
	}

	if len(query.Statements) != 2 {
		t.Fatalf("query.Statements does not contain 2 statements. got=%d",
			len(query.Statements))
	}

	tests := [][]string{
		{"sex"},
		{"name", "age"},
	}

	for i, tt := range tests {
		stmt := query.Statements[i]
		if !testDelStatement(t, stmt, tt) {
			return
		}
	}
}

func TestIntegerLiterals(t *testing.T) {
	input := `PUT num 42069;`

	l := lexer.New(input)
	p := New(l)
	query := p.ParseQuery()
	checkParserErrors(t, p)

	if len(query.Statements) != 1 {
		t.Fatalf("1 statement not found in query.")
	}
	stmt, ok := query.Statements[0].(*ast.PutStatement)
	if !ok {
		t.Fatalf("query.Statements[0] is not ast.PutStatement. got=%T",
			query.Statements[0])
	}

	checkIntegerLiteral(t, stmt.Value, 42069)
}

func TestStringLiterals(t *testing.T) {
	input := `PUT name "Manan"`

	l := lexer.New(input)
	p := New(l)
	query := p.ParseQuery()
	checkParserErrors(t, p)

	if len(query.Statements) != 1 {
		t.Fatalf("1 statement not found in query")
	}
	stmt, ok := query.Statements[0].(*ast.PutStatement)
	if !ok {
		t.Fatalf("query.Statements[0] is not ast.PutStatement. got=%T",
			query.Statements[0])
	}

	checkStringLiteral(t, stmt.Value, "Manan")
}

func TestBooleanLiteral(t *testing.T) {
	input := `PUT can_drink false;PUT can_marry true;`
	bools := []bool{false, true}

	l := lexer.New(input)
	p := New(l)
	query := p.ParseQuery()
	checkParserErrors(t, p)

	if len(query.Statements) != 2 {
		t.Fatalf("2 statements not found in query")
	}

	for i, stmt := range query.Statements {
		stmt, ok := stmt.(*ast.PutStatement)
		if !ok {
			t.Fatalf("query.Statements[0] is not ast.PutStatement. got=%T",
				query.Statements[0])
		}

		value, ok := stmt.Value.(*ast.BooleanLiteral)
		if !ok {
			t.Fatalf("value is not *ast.BooleanLiteral. got=%T", stmt.Value)
		}
		if value.Value != bools[i] {
			t.Fatalf("value is not %t. got=%T", bools[i], value.Value)
		}
	}
}

func TestParser_Errors(t *testing.T) {
	input := `GET PUT 31;`
	l := lexer.New(input)
	p := New(l)
	p.ParseQuery()
	if len(p.Errors()) == 0 {
		t.Error("no errors found. expected=2")
	}
}

func TestCollectionLiteral(t *testing.T) {
	input := `PUT ages [17, "Manan", ["nice"]];`

	l := lexer.New(input)
	p := New(l)
	query := p.ParseQuery()
	checkParserErrors(t, p)

	if len(query.Statements) != 1 {
		t.Fatalf("1 statement not found in query")
	}
	stmt, ok := query.Statements[0].(*ast.PutStatement)
	if !ok {
		t.Fatalf("query.Statements[0] is not ast.PutStatement. got=%T",
			query.Statements[0])
	}

	value, ok := stmt.Value.(*ast.CollectionLiteral)

	if !ok {
		t.Fatalf("value is not *ast.CollectionLiteral. got=%T", stmt.Value)
	}

	elements := value.Elements

	if len(elements) != 3 {
		t.Fatalf("expected 3 elements in collections. got=%d",
			len(elements))
	}

	checkIntegerLiteral(t, elements[0], 17)
	checkStringLiteral(t, elements[1], "Manan")

	collectionValue, ok := elements[2].(*ast.CollectionLiteral)
	collectionValueElements := collectionValue.Elements
	if !ok {
		t.Fatalf("value not *ast.CollectionLiteral. got=%T", elements[2])
	}
	if len(collectionValueElements) != 1 {
		t.Fatalf("expected 1 element in collection. got=%d", len(elements))
	}
	checkStringLiteral(t, collectionValueElements[0], "nice")

}

func testDelStatement(t *testing.T, s ast.Statement, keys []string) bool {
	if s.TokenLiteral() != "DEL" {
		t.Errorf("s.TokenLiteral not `del`. got=%q", s.TokenLiteral())
		return false
	}

	delStmt, ok := s.(*ast.DelStatement)
	if !ok {
		t.Errorf("s not *ast.DelStatement. got=%T", s)
		return false
	}

	if len(delStmt.Keys) != len(keys) {
		t.Errorf("len(delStmt.Keys) not %d. got=%d",
			len(keys), len(delStmt.Keys))
		return false
	}

	for i, key := range delStmt.Keys {
		if key.Value != keys[i] {
			t.Errorf("delStmt.Keys[%d] not '%s'. got=%s",
				i, keys[i], key.Value)
			return false
		}
	}
	return true
}

func testGetStatement(t *testing.T, s ast.Statement, keys []string) bool {

	if s.TokenLiteral() != "GET" {
		t.Errorf("s.TokenLiteral not `let`. got=%q", s.TokenLiteral())
		return false
	}

	getStmt, ok := s.(*ast.GetStatement)
	if !ok {
		t.Errorf("s not *ast.GetStatement. got=%T", s)
		return false
	}

	if len(getStmt.Keys) != len(keys) {
		t.Errorf("len(letStmt.Keys) not %d. got=%d",
			len(keys), len(getStmt.Keys))
		return false
	}

	for i, key := range getStmt.Keys {
		if key.Value != keys[i] {
			t.Errorf("letStmt.Keys[%d] not '%s'. got=%s",
				i, keys[i], key.Value)
			return false
		}
	}
	return true
}

func checkStringLiteral(t *testing.T, value ast.ValueExpression, toCompare string) {
	strValue, ok := value.(*ast.StringLiteral)
	if !ok {
		t.Fatalf("value not *ast.StringLiteral. got=%T", strValue.Value)
	}
	if strValue.Value != toCompare {
		t.Errorf("value.Value not %s. got=%s", toCompare, strValue.Value)
	}
	if value.TokenLiteral() != toCompare {
		t.Errorf("literal.TokenLiteral not %s. got=%s", toCompare,
			value.TokenLiteral())
	}
}

func checkIntegerLiteral(t *testing.T, value ast.ValueExpression, toCompare int64) {
	intVal, ok := value.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("value not *ast.IntegerLiteral. got=%T", value)
	}
	if intVal.Value != toCompare {
		t.Errorf("value.Value not %d. got=%d", 42069, intVal.Value)
	}
	if value.TokenLiteral() != strconv.Itoa(int(toCompare)) {
		t.Errorf("literal.TokenLiteral not %s. got=%s", "42069",
			value.TokenLiteral())
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
