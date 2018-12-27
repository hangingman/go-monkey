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

func (p *Parser) lookAheadIs(n int, t token.TokenType) bool {
	return p.lookAhead(n).Type == t
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
		fmt.Printf("ParseProgram: %s\n", stmt)
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.curIndex++
	}

	return program
}

func (p *Parser) parseStatement() ast.Statement {
	if p.curToken().Type == token.VAR {
		return p.parseVarStatement()
	}
	if p.isSimpleStatement() {
		// 単純文
		return p.parseSimpleStatement()
	}
	// 複合文
	return p.parseCompoundStatement()
}

// parseVarStatement は以下のような構文を解析する
// <変数宣言>::=var<変数名>{","<変数名>}";"
func (p *Parser) parseVarStatement() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curToken()}

	if !p.expectPeek(token.IDENT) {
		return nil
	}

	for {
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

func (p *Parser) isSimpleStatement() bool {

	if p.curTokenIs(token.INPUT) || p.curTokenIs(token.OUTPUT) {
		if p.lookAheadIs(2, token.SEMICOLON) {
			return false
		}
	}
	// fmt.Printf("isSimpleStatement?: %t\n", true)
	return true
}

func (p *Parser) parseSimpleStatement() *ast.SimpleStatement {
	stmt := &ast.SimpleStatement{Token: p.curToken()}

	switch p.curToken().Type {
	case token.INPUT:
		stmt.Name = ast.Identifier{
			Token: p.peekToken(),
			Value: p.peekToken().Literal,
		}
		p.curIndex++
	case token.OUTPUT:
		stmt.Name = ast.Identifier{
			Token: p.peekToken(),
			Value: p.peekToken().Literal,
		}
		p.curIndex++
	default:
		return nil
	}

	return stmt
}

func (p *Parser) parseCompoundStatement() *ast.CompoundStatement {
	stmt := &ast.CompoundStatement{Tokens: []token.Token{}}

	for {
		switch p.curToken().Type {
		case token.INPUT:
			simpleStmt := p.parseSimpleStatement()
			p.curIndex++
			stmt.Statements = append(stmt.Statements, *simpleStmt)
			stmt.Tokens = append(stmt.Tokens, simpleStmt.Token)
		case token.OUTPUT:
			simpleStmt := p.parseSimpleStatement()
			p.curIndex++
			stmt.Statements = append(stmt.Statements, *simpleStmt)
			stmt.Tokens = append(stmt.Tokens, simpleStmt.Token)
		default:
			return nil
		}
		if p.peekToken().Type == token.EOF {
			break
		}
		p.curIndex++
	}

	return stmt
}
