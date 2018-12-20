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

type SimpleStatement struct {
	Token              token.Token
	Name               Identifier
	Expression         Expression
	CompoundStatements []CompoundStatement
}

func (ss *SimpleStatement) statementNode()       {}
func (ss *SimpleStatement) TokenLiteral() string { return ss.Token.Literal }

// CompoundStatement は <複合文>::=<単純文>";"<単純文>{";"<単純文>} を表す
type CompoundStatement struct {
	Tokens     []token.Token
	Statements []SimpleStatement
}

func (cs *CompoundStatement) statementNode()       {}
func (cs *CompoundStatement) TokenLiteral() string { return cs.Tokens[0].Literal }

// VarStatement は`var a = 10;` のような構文を解析する
type VarStatement struct {
	Token token.Token
	Names []Identifier
}

func (vs *VarStatement) statementNode()       {}
func (vs *VarStatement) TokenLiteral() string { return vs.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }
