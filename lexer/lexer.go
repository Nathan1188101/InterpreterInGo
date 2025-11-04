package lexer //will be part of the lexer package

import "monkey/token"

type Lexer struct {
	input        string //
	position     int    //32 bits, current position in input (points to current char)
	readPosition int    //32 bits, current reading position in input (after current char)
	ch           byte   //8  bits, current char under examination
}

//takes in input (which is a string), the input is the source code/text lexer will process
//the function returns a reference to a lexer
func New(input string) *Lexer {
	l := &Lexer{input: input} //allocating memory for a new instance the lexer (allocated with &Lexer{})
	l.readChar()//starts the lexer off (called on the newly created lexer "l" instance)
	return l
}

func (l *Lexer) readChar() {
	//first checking if we are at the end of the input (base case)
	if l.readPosition >= len(l.input) {
		l.ch=0 //if it is the end of input, set l.ch to 0 which is the ASCII code for "nul"
	}
	//if not at the end of input, set l.ch to next char
	else {
		//do this by accessing
		l.ch = l.input[l.readPosition]
	}
	//readPosition points to the "next" character in the input, position points to the character in the input that corresponds to the ch byte 
	l.position = l.readPosition //moves to current
	l.readPosition += 1 //read position moves on (incremented 1 for next char)
}

