package main

import (
	"fmt"
	"os"
)

type TokenType int

const (
	LeftParen TokenType = iota
	RightParen
	LeftBrace
	RightBrace

	Comma
	Dot
	Minus
	Plus
	Semicolon
	Slash
	Star

	// One or two character tokens.
	Bang
	BangEqual
	Equal
	EqualEqual
	Greater
	GreaterEqual
	Less
	LessEqual

	// Literals.
	Identifier
	String
	Number

	// Keywords.
	And
	Class
	Else
	False
	Fun
	For
	If
	Nil
	Or
	Print
	Return
	Super
	This
	True
	Var
	While

	Eof
)

var Tokens = [...]string{
	"LeftParen",
	"RightParen",
	"LeftBrace",
	"RightBrace",

	"Comma",
	"Dot",
	"Minus",
	"Plus",
	"Semicolon",
	"Slash",
	"Star",

	// One or two character tokens.
	"Bang",
	"BangEqual",
	"Equal",
	"EqualEqual",
	"Greater",
	"GreaterEqual",
	"Less",
	"LessEqual",

	// Literals.
	"Identifier",
	"String",
	"Number",

	// Keywords.
	"And",
	"Class",
	"Else",
	"False",
	"Fun",
	"For",
	"If",
	"Nil",
	"Or",
	"Print",
	"Return",
	"Super",
	"This",
	"True",
	"Var",
	"While",

	"Eof",
}

func (t TokenType) String() string {
	if t >= LeftParen && t <= Eof {
		return Tokens[t]
	}

	fmt.Fprintln(os.Stderr, "Token type is unknown.")

	return ""
}

type LoxToken struct {
	Type    TokenType
	Lexeme  string
	Literal interface{} // This is the equivalent of Java's Object type
	Line    int64
}

func NewLoxToken(ttype TokenType, lex string, literal interface{}, line int) *LoxToken {
	return &LoxToken{
		Type:    ttype,
		Lexeme:  lex,
		Literal: literal,
		Line:    line,
	}
}

func (t *LoxToken) String() string {
	return fmt.Sprintf("%s %s %v\n", t.Type.String(), t.Lexeme, t.Literal)
}
