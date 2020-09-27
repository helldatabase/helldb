package evaluator

import (
	"bufio"
	"fmt"
	"helldb/engine/types"
	"os"

	s "helldb/engine/store"
	"helldb/query/ast"
	"helldb/query/lexer"
	"helldb/query/parser"
)

var Store = s.Init()

func REPL(prompt string) {
	fmt.Println("helldb client v0.0.1")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt)
	for scanner.Scan() {
		input := scanner.Text()
		if input == "dumb" {
			fmt.Println(Store.JSON())
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
		results := make([][]types.BaseType, len(query.Statements))
		for i, statement := range query.Statements {
			if valid, isGet := isGetOrDelStatement(statement); valid {
				var keys []string
				if isGet {
					keys = keysFromGetStatement(statement)
					results[i] = Store.Get(keys)
				} else {
					keys = keysFromDelStatement(statement)
					results[i] = Store.Del(keys)
				}
			} else {
				putStatement := statement.(*ast.PutStatement)
				key, value := putStatement.Key.String(), putStatement.Value
				Store.Put(key, ast.ExtractToBaseType(value))
				results[i] = nil
			}
		}
		return Response{Errors: p.Errors(), Results: results}
	}
	return Response{Errors: p.Errors(), Results: nil}
}
