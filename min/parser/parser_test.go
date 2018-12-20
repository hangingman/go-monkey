package parser

import (
	"fmt"
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
	"github.com/hangingman/go-monkey/min/token"
	"testing"
)

func TestParserNew(t *testing.T) {
	input := `
var x;
var y;
var foo, bar;
var i, j, k;
`
	l := lexer.New(input)
	p := New(l)

	actual := len(p.lexedTokens)
	expected := 19
	if actual != expected {
		t.Errorf("lexedTokens: got %v, want %v", actual, expected)
	}
	actualTok1 := p.lookAhead(1)
	expectTok1 := token.Token{Type: token.IDENT, Literal: "x"}
	if actualTok1 != expectTok1 {
		t.Errorf("lookAhead(1): got %v, want %v", actualTok1, expectTok1)
	}
	actualTok2 := p.lookAhead(2)
	expectTok2 := token.Token{Type: token.SEMICOLON, Literal: ";"}
	if actualTok2 != expectTok2 {
		t.Errorf("lookAhead(2): got %v, want %v", actualTok2, expectTok2)
	}
	actualTok3 := p.lookAhead(3)
	expectTok3 := token.Token{Type: token.VAR, Literal: "var"}
	if actualTok3 != expectTok3 {
		t.Errorf("lookAhead(3): got %v, want %v", actualTok3, expectTok3)
	}
}

func TestVarStatements(t *testing.T) {
	input := `
var x;
var y;
var foo, bar;
var i, j, k;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}
	if len(program.Statements) != 4 {
		t.Fatalf("program.Statements does not contain 4 statements. got=%d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifer []string
	}{
		{[]string{"x"}},
		{[]string{"y"}},
		{[]string{"foo", "bar"}},
		{[]string{"i", "j", "k"}},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testVarStatement(t, stmt, tt.expectedIdentifer) {
			return
		}
	}
}

func testVarStatement(t *testing.T, s ast.Statement, expectNames []string) bool {

	if s.TokenLiteral() != "var" {
		t.Errorf("s.TokenLiteral not 'var'. got=%q", s.TokenLiteral())
		return false
	}

	varStmt, ok := s.(*ast.VarStatement)
	if !ok {
		t.Errorf("s not *ast.VarStatement. got=%T", s)
		return false
	}

	if len(expectNames) != len(varStmt.Names) {
		// fmt.Printf("%s", varStmt.Names)
		t.Errorf("varStmt.Names size not '%d', got=%d", len(expectNames), len(varStmt.Names))
		return false
	}

	for i, expectName := range expectNames {
		if expectName != varStmt.Names[i].Value {
			t.Errorf("varStmt.Names[%d].Value not '%s', got=%s", i, expectName, varStmt.Names[i].Value)
			return false
		}
	}

	return true
}

func TestSimpleStatements(t *testing.T) {
	input := `input x`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	// checkParserErrors(t, p)
	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	for stmt := range program.Statements {
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
