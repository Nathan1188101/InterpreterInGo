package lexer //will be part of the lexer package

type Lexer struct {
	input        string //
	position     int    //32 bits, current position in input (points to current char)
	readPosition int    //32 bits, current reading position in input (after current char)
	ch           byte   //8  bits, current char under examination
}

//returns *Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) readChar() {
	//first checking if we are at the end of the input (base case)
	if l.readPosition >= len(l.input) {
		l.ch=0 //if it is the end of input, set l.ch to 0 which is the ASCII code for "nul"
	}
	else{
		l.ch = l.input[l.readPosition]
	}
	//readPosition points to the "next" character in the input, position points to the character in the input that corresponds to the ch byte 
	l.position = l.readPosition
	l.readPosition += 1
}

