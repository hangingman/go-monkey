package lexer

import (
	"github.com/hangingman/go-monkey/min/token"
	"testing"
)

func TestStatements(t *testing.T) {
	input := `var x, i;
  input x;
  if x=0 then x:=1 fi;
  i:=x-1;
  while i>=2
    begin
      x:=x*i;
      i:=i-1
    end;

  output x;
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.VAR, "var"},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "i"},
		{token.SEMICOLON, ";"},
		{token.INPUT, "input"},
		{token.IDENT, "x"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.IDENT, "x"},
		{token.ASSIGN, "="},
		{token.INT, "0"},
		{token.THEN, "then"},
		{token.IDENT, "x"},
		{token.COLONEQ, ":="},
		{token.INT, "1"},
		{token.FI, "fi"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.COLONEQ, ":="},
		{token.IDENT, "x"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.WHILE, "while"},
		{token.IDENT, "i"},
		{token.GTEQ, ">="},
		{token.INT, "2"},
		{token.BEGIN, "begin"},
		{token.IDENT, "x"},
		{token.COLONEQ, ":="},
		{token.IDENT, "x"},
		{token.ASTERISK, "*"},
		{token.IDENT, "i"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "i"},
		{token.COLONEQ, ":="},
		{token.IDENT, "i"},
		{token.MINUS, "-"},
		{token.INT, "1"},
		{token.END, "end"},
		{token.SEMICOLON, ";"},
		{token.OUTPUT, "output"},
		{token.IDENT, "x"},
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
