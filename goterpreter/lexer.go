package main

type Expr interface {
	Evaluate() interface{}
	Visit() string
}

type Binary struct {
	left  *Expr
	op    LoxToken
	right *Expr
}

func (b *Binary) Evaluate() {
	return b.left.Evaluate()
}

func (b *Binary) Visit() string {
	return fmt.Sprintf("(%s %s %s)", left.Visit(), op.String(), right.Visit())
}

type Unary struct {
	op    LoxToken
	right *Expr
}

func (u *Unary) Evaluate() {
	return b.right.Evaluate()
}

func (u *Unary) Visit() string {
	return fmt.Sprintf("(%s %s)", op.String(), right.Visit())
}

type Literal struct {
	value interface{}
}

func (l *Literal) Evaluate() {
	return value
}

func (l *Literal) Visit() string {
	return fmt.Sprintf("%v", l.value)
}

type Grouping struct {
	expression *Expr
}

func (g *Grouping) Evaluate() {
	return expression.Evaluate()
}

func (g *Grouping) Visit() string {
	return fmt.Sprintf("%s", expr.Visit())
}
