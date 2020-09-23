package parser

import (
	"fmt"
	"strconv"

	"helldb/query/ast"
	"helldb/query/token"
)

func (p *Parser) parseIntegerLiteral() ast.ValueExpression {
	lit := &ast.IntegerLiteral{Token: p.curToken}
	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (p *Parser) parseKeys() []ast.Identifier {
	var keys []ast.Identifier
	ident := ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	keys = append(keys, ident)
	for p.peekTokenIs(token.AND) {
		p.nextToken()
		p.nextToken()
		ident = ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		keys = append(keys, ident)
	}
	if !p.expectPeek(token.SEMICOLON) {
		return nil
	}
	return keys
}

func (p *Parser) parseCollectionLiteral() ast.ValueExpression {
	collection := &ast.CollectionLiteral{Token: p.curToken}
	collection.Elements = p.parseExpressionCollectionElements(token.RBRACK)
	return collection
}

func (p *Parser) parseStringLiteral() ast.ValueExpression {
	lit := &ast.StringLiteral{Token: p.curToken}
	lit.Value = p.curToken.Literal
	return lit
}

func (p *Parser) parseBooleanLiteral() ast.ValueExpression {
	lit := &ast.BooleanLiteral{Token: p.curToken}
	lit.Value = p.curToken.Literal == "true"
	return lit
}

func (p *Parser) parseExpressionCollectionElements(end token.Type) []ast.ValueExpression {

	var list []ast.ValueExpression

	if p.peekTokenIs(end) {
		p.nextToken()
		return list
	}
	p.nextToken()
	list = append(list, p.parseExpression())
	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		list = append(list, p.parseExpression())
	}

	if !p.expectPeek(end) {
		return nil
	}

	return list
}
