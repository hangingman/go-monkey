package lexer

import (
	"testing"
	"bitbucket.org/hangedman/go-monkey/token"
)

func TestStatements(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
  x + y;
}

let result = add(five, ten);
`

}

func TestNextToken(t *testing.T) {
	input := `=+(),;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)            
		}
	}
}
