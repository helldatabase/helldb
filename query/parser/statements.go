package parser

import (
	"helldb/query/ast"
	"helldb/query/token"
)

func (p *Parser) parsePutStatement() *ast.PutStatement {
	stmt := &ast.PutStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	} else {
		stmt.Key = ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		p.nextToken()
		stmt.Value = p.parseExpression()
		return stmt
	}
}

func (p *Parser) parseGetStatement() *ast.GetStatement {
	stmt := &ast.GetStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	} else {
		stmt.Keys = p.parseKeys() // make([]ast.Identifier, 0, 1)
		return stmt
	}
}

func (p *Parser) parseDelStatement() *ast.DelStatement {
	stmt := &ast.DelStatement{Token: p.curToken}
	if !p.expectPeek(token.IDENT) {
		return nil
	} else {
		stmt.Keys = p.parseKeys()
		return stmt
	}
}
