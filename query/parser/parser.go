package parser

import (
	"helldb/query/ast"
	"helldb/query/lexer"
	"helldb/query/token"
)

type valueParseFn func() ast.ValueExpression

type Parser struct {
	l *lexer.Lexer

	errors []string

	curToken  token.Token
	peekToken token.Token

	valueParseFns map[token.Type]valueParseFn
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.nextToken()
	p.nextToken()

	p.valueParseFns = make(map[token.Type]valueParseFn)

	p.registerValueParseFn(token.INT, p.parseIntegerLiteral)
	p.registerValueParseFn(token.BOOL, p.parseBooleanLiteral)
	p.registerValueParseFn(token.STRING, p.parseStringLiteral)
	p.registerValueParseFn(token.LBRACK, p.parseCollectionLiteral)

	return p
}

func (p *Parser) ParseQuery() *ast.Query {
	query := &ast.Query{}
	query.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			query.Statements = append(query.Statements, stmt)
		}
		p.nextToken()
	}
	return query
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.GET:
		return p.parseGetStatement()
	case token.PUT:
		return p.parsePutStatement()
	case token.DEL:
		return p.parseDelStatement()
	default:
		return nil
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) parseExpression() ast.ValueExpression {
	fn := p.valueParseFns[p.curToken.Type]
	if fn == nil {
		return nil
	}
	value := fn() // parse value using function
	return value
}
