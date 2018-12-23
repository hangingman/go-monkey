package ast

import "github.com/hangingman/go-monkey/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// LetStatement は`let a = 10;` のような構文を解析する
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

// ReturnStatement は`return 10;` のような構文を解析する
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ls *ReturnStatement) statementNode()       {}
func (ls *ReturnStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// ExpressionStatement は全ての式の解析に使われる文
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode()       {}
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }
