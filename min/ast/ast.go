package ast

import "github.com/hangingman/go-monkey/min/token"

type Node interface {
	TokenLiteral() string
}

// Statement (文)
type Statement interface {
	Node
	statementNode()
}

// Expression (式)
type Expression interface {
	Node
	expressionNode()
}

// Program (プログラム)
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

// VarStatement は`var a = 10;` のような構文を解析する
type VarStatement struct {
	Token token.Token
	Names []Identifier
}

func (ls *VarStatement) statementNode()       {}
func (ls *VarStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
