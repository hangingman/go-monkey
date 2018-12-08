package token

// TokenType は stringのエイリアス
type TokenType string

// Token は文字列をトークン化した後の情報を保持する
// TokenType にはtoken内で定義したconst値
// Literal には実際に読み取った文字列
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "ILLEGAL"
	EOF       = "EOF"
	IDENT     = "IDENT"
	INT       = "INT"
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	ASTERISK  = "*"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LT        = "<"
	GT        = ">"
	LTEQ      = "<="
	GTEQ      = ">="
	VAR       = "VAR"
	TRUE      = "TRUE"
	FALSE     = "FALSE"
	IF        = "IF"
	THEN      = "THEN"
	ELSE      = "ELSE"
	FI        = "FI"
	WHILE     = "WHILE"
	BEGIN     = "BEGIN"
	END       = "END"
	INPUT     = "INPUT"
	OUTPUT    = "OUTPUT"
	COLONEQ   = ":="
	EQ        = "=="
	NOTEQ     = "!="
)

var keywords = map[string]TokenType{
	"var":    VAR,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"then":   THEN,
	"else":   ELSE,
	"fi":     FI,
	"while":  WHILE,
	"begin":  BEGIN,
	"end":    END,
	"input":  INPUT,
	"output": OUTPUT,
}

func LookupIndent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
