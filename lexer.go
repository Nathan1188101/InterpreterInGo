package lexer //will be part of the lexer package

type Lexer struct {
	input        string //
	position     int    //32 bits, current position in input (points to current char)
	readPosition int    //32 bits,currewnt reading position in input (after current char)
	ch           byte   //8  bits, current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}
