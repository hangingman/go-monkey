package parser

import (
	"fmt"
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
	"github.com/hangingman/go-monkey/min/token"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
	errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	// ２つトークンを読み込む。curTokenとpeekTokenの両方がセットされる。
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
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

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.VAR:
		return p.parseVarStatement()
	default:
		return nil
	}
}

// parseVarStatement は以下のような構文を解析する
// <変数宣言>::=var<変数名>{","<変数名>}";"
func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	for {
		// fmt.Printf("cur: %s, peek: %s\n", p.curToken, p.peekToken)
		stmt.Names = append(stmt.Names, ast.Identifier{Token: p.curToken, Value: p.curToken.Literal})
		if !p.expectPeek(token.COMMA) || p.expectPeek(token.SEMICOLON) {
			break
		}
		p.nextToken()
	}

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}
