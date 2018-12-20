package parser

import (
	"fmt"
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
	"github.com/hangingman/go-monkey/min/token"
)

type Parser struct {
	l           *lexer.Lexer
	lexedTokens []*token.Token
	curIndex    int
	errors      []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:           l,
		lexedTokens: []*token.Token{},
		curIndex:    0,
		errors:      []string{},
	}
	// EOFまでトークンを読み込む
	for {
		tok := p.l.NextToken()
		p.lexedTokens = append(p.lexedTokens, &tok)
		if tok.Type == token.EOF {
			break
		}
	}
	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) curToken() token.Token {
	return *p.lexedTokens[p.curIndex]
}

func (p *Parser) peekToken() token.Token {
	return *p.lexedTokens[p.curIndex+1]
}

func (p *Parser) lookAhead(n int) token.Token {
	if p.curIndex+n < len(p.lexedTokens) {
		return *p.lexedTokens[p.curIndex+n]
	}
	return token.Token{Type: token.ILLEGAL, Literal: ""}
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken().Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken().Type == t
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken().Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.curIndex++
		return true
	}
	p.peekError(t)
	return false
}

// ParseProgram は Parser を受け取ってAST化されたProgramを返す
func (p *Parser) ParseProgram() *ast.Program {
	// 中括弧{}は配列の宣言を表す
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken().Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.curIndex++
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken().Type {
	case token.VAR:
		return p.parseVarStatement()
	default:
		return p.parseCompoundStatement()
	}
}

// parseVarStatement は以下のような構文を解析する
// <変数宣言>::=var<変数名>{","<変数名>}";"
func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken()}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	for {
		// fmt.Printf("cur: %s, peek: %s\n", p.curToken, p.peekToken)
		stmt.Names = append(stmt.Names, ast.Identifier{Token: p.curToken(), Value: p.curToken().Literal})
		if !p.expectPeek(token.COMMA) || p.expectPeek(token.SEMICOLON) {
			break
		}
		p.curIndex++
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.curIndex++
	}

	return stmt
}

func (p *Parser) parseCompoundStatement() *ast.CompoundStatement {
	stmt := &ast.CompoundStatement{}

	switch p.curToken().Type {
	case token.INPUT:
	case token.OUTPUT:
		if !p.expectPeek(token.IDENT) {
			return nil
		}
	case token.IF:
	case token.WHILE:
	case token.IDENT:
	}

	return stmt
}
