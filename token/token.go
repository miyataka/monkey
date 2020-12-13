package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// indentifier + literal
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 134356

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// delimeter
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keyword
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
