package evaluator

import (
	"encoding/json"
	"fmt"
	"helldb/engine/types"
	"helldb/query/ast"
	"helldb/query/token"
)

func showResults(results []types.BaseType) {
	for _, result := range results {
		var res interface{}
		if result != nil {
			res = result.Native()
		} else {
			res = nil
		}
		str, _ := json.MarshalIndent(res, "", "  ")
		fmt.Println(string(str))
	}
}

func showErrors(errors []string) {
	fmt.Println("errors: ")
	for i, errorMsg := range errors {
		fmt.Println(i, errorMsg)
	}
}

func isGetOrDelStatement(statement ast.Statement) (bool, bool) {
	if statement.TokenLiteral() == token.GET {
		return true, true
	} else if statement.TokenLiteral() == token.DEL {
		return true, false
	} else {
		return false, false
	}
}

func keysFromGetStatement(statement ast.Statement) []string {
	getStatement := statement.(*ast.GetStatement)
	keys := make([]string, len(getStatement.Keys))
	for i, key := range getStatement.Keys {
		keys[i] = key.Value
	}
	return keys
}

func keysFromDelStatement(statement ast.Statement) []string {
	delStatement := statement.(*ast.DelStatement)
	keys := make([]string, len(delStatement.Keys))
	for i, key := range delStatement.Keys {
		keys[i] = key.Value
	}
	return keys
}
