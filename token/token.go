package token

type TokenType string

// defining the structure of a token (fundamental to interpreter)
type Token struct {

	// each token has:
	// Type of TokenType -> sounds confusing but we are just categorizing the token type, i.e IDENT, INT, EOF, PLUS, etc.
	// Literal of type string -> holds the actual value or character from the source code

	//exmaples:
	// if the token is a INT -> the literal may be "123"

	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL" //signifies a token/character we don't know about
	EOF     = "EOF"     //end of file

	//identifiers + literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"   //1234567

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
