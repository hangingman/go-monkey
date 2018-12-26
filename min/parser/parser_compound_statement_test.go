package parser

import (
	"fmt"
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
	"github.com/hangingman/go-monkey/min/token"
	"testing"
)

func TestCompoundStatements(t *testing.T) {
	input := `input x; output y;`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	for _, stmt := range program.Statements {
		fmt.Printf("%s\n", stmt)
	}

	// if len(program.Statements) != 1 {
	//     t.Fatalf("program.Statements does not contain 1 statements. got=%d",
	//     	len(program.Statements))
	// }

	// tests := []struct {
	// 	expectedToken     token.Token
	// 	expectedIdentifer string
	// }{
	// 	{token.Token{Type: token.INPUT, Literal: string("input")}, "x"},
	// 	{token.Token{Type: token.OUTPUT, Literal: string("output")}, "y"},
	// }

	// for i, tt := range tests {
	// 	stmt := program.Statements[i]
	// 	if !testCompoundStatement(t, stmt, tt.expectedToken, tt.expectedIdentifer) {
	// 	    return
	// 	}
	// }
}

func testCompoundStatement(t *testing.T, s ast.Statement, expectedToken token.Token, expectedName string) bool {

	// compStmt, ok := s.(*ast.CompoundStatement)
	// if !ok {
	// 	t.Errorf("s not *ast.CompoundStatement. got=%T", s)
	// 	return false
	// }

	// if expectedToken != compStmt.Token {
	// 	t.Errorf("compStmt.Name.Token not '%s', got=%s", expectedToken, compStmt.Token)
	// 	return false
	// }

	// if expectedName != compStmt.Name.Value {
	// 	t.Errorf("compStmt.Name.Value not '%s', got=%s", expectedName, compStmt.Name.Value)
	// 	return false
	// }

	return true
}
