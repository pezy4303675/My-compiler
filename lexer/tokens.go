package lexer

type TokenType string

const (
	TOKEN_LET   TokenType = "LET"
	TOKEN_IDENT TokenType = "IDENT"
	TOKEN_INT   TokenType = "INT"
	TOKEN_EQ    TokenType = "="
	TOKEN_PLUS  TokenType = "+"
	TOKEN_SEMI  TokenType = ";"
)

type Token struct {
	Type    TokenType
	Literal string
}
