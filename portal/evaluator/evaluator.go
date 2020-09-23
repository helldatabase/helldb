package evaluator

import (
	"bufio"
	"fmt"
	"os"

	s "helldb/engine/store"
	"helldb/query/ast"
	"helldb/query/lexer"
	"helldb/query/parser"
)

var store = s.Init()

func REPL(prompt string) {
	fmt.Println("helldb client v0.0.1")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "json" {
			fmt.Println(store.JSON())
			fmt.Print(prompt)
			continue
		}
		resp := Eval(input)
		if len(resp.Errors) != 0 {
			showErrors(resp.Errors)
		} else {
			showResults(resp.Results)
		}
		fmt.Print(prompt)
	}
}

func Eval(input string) Response {
	l := lexer.New(input)
	p := parser.New(l)
	query := p.ParseQuery()
	if len(p.Errors()) == 0 {
		for _, statement := range query.Statements {
			if valid, isGet := isGetOrDelStatement(statement); valid {
				var keys []string
				if isGet {
					keys = keysFromGetStatement(statement)
				} else {
					keys = keysFromDelStatement(statement)
				}
				return Response{Errors: nil, Results: store.Get(keys)}
			} else {
				putStatement := statement.(*ast.PutStatement)
				key := putStatement.Key.String()
				element := putStatement.Value
				switch element.(type) {
				case *ast.IntegerLiteral:
					store.Put(key, element.(*ast.IntegerLiteral).ToBaseType())
				case *ast.StringLiteral:
					store.Put(key, element.(*ast.StringLiteral).ToBaseType())
				case *ast.BooleanLiteral:
					store.Put(key, element.(*ast.BooleanLiteral).ToBaseType())
				case *ast.CollectionLiteral:
					store.Put(key, element.(*ast.CollectionLiteral).ToBaseType())
				}
				return Response{Errors: nil, Results: nil}
			}
		}
	}
	return Response{Errors: p.Errors(), Results: nil}
}
