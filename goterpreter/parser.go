package main

type Parser struct {
	tokens  []LoxToken
	current int64
}

func NewParser(tokens []LoxToken) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) expression() *Expr {
	return p.equality()
}

func (p *Parser) equality() *Expr {
	expr := comparison()

	for {
		if !match(BangEqual, EqualEqual) {
			break
		}
		operator := previous()
		right := comparison()
		expr = Binary{left: expr, op: operator, right: right}
	}

	return expr
}

func (p *Parser) match(types []LoxToken) bool {
	for _, t := range types {
		if check(t) {
			advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(TokenType t) bool {
	if p.isAtEnd() {
		return false
	}

	return t == peek()
}

func (p *Parser) advance() LoxToken {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return peek() == EOF
}

func (p *Parser) peek() LoxToken {
	return p.tokens[p.current]
}

func (p *Parser) previous() LoxToken {
	return p.tokens[p.current-1]
}
