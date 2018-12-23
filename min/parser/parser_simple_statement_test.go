package parser

import (
	"fmt"
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
	"testing"
)

func TestSimpleStatements(t *testing.T) {
	input := `input x`

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

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer string
	}{
		{"x"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSimpleStatement(t, stmt, tt.expectedIdentifer) {
			return
		}
	}
}

func testSimpleStatement(t *testing.T, s ast.Statement, expectNames string) bool {

	if s.TokenLiteral() != "input" {
		t.Errorf("s.TokenLiteral not 'var'. got=%q", s.TokenLiteral())
		return false
	}

	// simpleStmt, ok := s.(*ast.SimpleStatement)
	// if !ok {
	// 	t.Errorf("s not *ast.SimpleStatement. got=%T", s)
	// 	return false
	// }

	// if len(expectNames) != len(simpleStmt.Names) {
	// 	// fmt.Printf("%s", simpleStmt.Names)
	// 	t.Errorf("simpleStmt.Names size not '%d', got=%d", len(expectNames), len(simpleStmt.Names))
	// 	return false
	// }

	// for i, expectName := range expectNames {
	// 	if expectName != simpleStmt.Names[i].Value {
	// 		t.Errorf("simpleStmt.Names[%d].Value not '%s', got=%s", i, expectName, simpleStmt.Names[i].Value)
	// 		return false
	// 	}
	// }

	return true
}
