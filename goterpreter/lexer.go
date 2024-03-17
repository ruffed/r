package main

type Expr struct {
	left  *Expr
	op    LoxToken
	right *Expr
}
