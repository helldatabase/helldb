package main

import (
	"fmt"
	"helldb/query/lexer"
	"helldb/query/token"
)

func main() {
	l := lexer.New(`PUT name "Manan";`)
	for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
		fmt.Printf("%v\n", t)
	}
}
