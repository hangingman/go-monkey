package parser

import (
	"github.com/hangingman/go-monkey/min/ast"
	"github.com/hangingman/go-monkey/min/lexer"
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
