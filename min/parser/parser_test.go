package parser

import (
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

func TestIsSimpleStatement(t *testing.T) {
	inputs := []string{
		"input x",
		"input x;",
	}

	tests := []struct {
		expectedAnswer bool
	}{
		{true},
		{false},
	}

	for i, input := range inputs {
		l := lexer.New(input)
		p := New(l)
		actualAnswer := p.isSimpleStatement()
		if tests[i].expectedAnswer != actualAnswer {
			t.Errorf("isSimpleStatement(%s) got %v, want %v", input, actualAnswer, tests[i].expectedAnswer)
		}
	}
}
