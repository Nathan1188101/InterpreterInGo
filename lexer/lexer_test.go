package lexer //we are creating the package, we are declaring that this file belongs to the "lexer" package

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {

	input := `=+(){},;` //:= infers the type (string in this case) (initializes and declares)

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		//test cases. "Expected test data" we compare it against the actual output of the lexer (I think)
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}

}
