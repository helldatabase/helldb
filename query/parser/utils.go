package parser

import (
	"fmt"

	"helldb/query/token"
)

func (p *Parser) registerValueParseFn(tokenType token.Type, fn valueParseFn) {
	p.valueParseFns[tokenType] = fn
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) peekError(t token.Type) {
	msg := fmt.Sprintf("expected next token to be %s, got %s (%s) instead",
		t, p.peekToken.Type, p.peekToken.Literal)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}
