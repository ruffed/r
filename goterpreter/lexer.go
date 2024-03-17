package main

import "fmt"

type Expr interface {
	Evaluate() interface{}
	Visit() string
}

type Binary struct {
	left  Expr
	op    LoxToken
	right Expr
}

func (b *Binary) Evaluate() interface{} {
	return b.left.Evaluate()
}

func (b *Binary) Visit() string {
	return fmt.Sprintf("(%s %s %s)", b.left.Visit(), b.op.String(), b.right.Visit())
}

type Unary struct {
	op    LoxToken
	right Expr
}

func (u *Unary) Evaluate() interface{} {
	return u.right.Evaluate()
}

func (u *Unary) Visit() string {
	return fmt.Sprintf("(%s %s)", u.op.String(), u.right.Visit())
}

type Literal struct {
	value interface{}
}

func (l *Literal) Evaluate() interface{} {
	return l.value
}

func (l *Literal) Visit() string {
	return fmt.Sprintf("%v", l.value)
}

type Grouping struct {
	expression Expr
}

func (g *Grouping) Evaluate() interface{} {
	return g.expression.Evaluate()
}

func (g *Grouping) Visit() string {
	return fmt.Sprintf("%s", g.expression.Visit())
}
