package lexer

import (
	"github.com/hangingman/go-monkey/rcf/token"
	"testing"
)

func TestStatements(t *testing.T) {
	input := `def fact(x) = if x>=2 then fact(x-1)*x else 1 fi
in fact(3)
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.DEF, "def"},
		{token.IDENT, "fact"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.ASSIGN, "="},
		{token.IF, "if"},
		{token.IDENT, "x"},
		{token.GTEQ, ">="},
		{token.INT, "2"},
		{token.THEN, "then"},
		{token.IDENT, "fact"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.RPAREN, ")"},
		{token.ASTERISK, "*"},
		{token.IDENT, "x"},
		{token.ELSE, "else"},
		{token.INT, "1"},
		{token.FI, "fi"},
		{token.IN, "in"},
		{token.IDENT, "fact"},
		{token.LPAREN, "("},
		{token.INT, "3"},
		{token.RPAREN, ")"},
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
