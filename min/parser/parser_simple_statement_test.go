package parser

import (
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
	"github.com/hangingman/go-monkey/min/token"    
	"testing"
)

func TestSimpleStatements(t *testing.T) {
	input := `input x
output y
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	// for _, stmt := range program.Statements {
	// 	fmt.Printf("%s\n", stmt)
	// }

	if len(program.Statements) != 2 {
		t.Fatalf("program.Statements does not contain 2 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
        expectedToken token.Token
		expectedIdentifer string
	}{
		{token.Token{Type: token.INPUT, Literal: string("input")}, "x"},
        {token.Token{Type: token.OUTPUT, Literal: string("output")}, "y"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testSimpleStatement(t, stmt, tt.expectedToken, tt.expectedIdentifer) {
			return
		}
	}
}

func testSimpleStatement(t *testing.T, s ast.Statement, expectedToken token.Token, expectedName string) bool {

	simpleStmt, ok := s.(*ast.SimpleStatement)
	if !ok {
		t.Errorf("s not *ast.SimpleStatement. got=%T", s)
		return false
	}

    if expectedToken != simpleStmt.Token {
        t.Errorf("simpleStmt.Name.Token not '%s', got=%s", expectedToken, simpleStmt.Token)
        return false
    }
    
    if expectedName != simpleStmt.Name.Value {
        t.Errorf("simpleStmt.Name.Value not '%s', got=%s", expectedName, simpleStmt.Name.Value)
        return false
    }

	return true
}
