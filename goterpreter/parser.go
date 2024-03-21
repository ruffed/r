package main

type Parser struct {
	tokens  []LoxToken
	current int64
}

func NewParser(tokens []LoxToken) *Parser {
	return &Parser{tokens, 0}
}

func (p *Parser) parse() Expr {
	return p.expression()
}

func (p *Parser) expression() Expr {
	return p.equality()
}

func (p *Parser) equality() Expr {
	expr := p.comparison()

	for {
		if !p.match([]TokenType{BangEqual, EqualEqual}) {
			break
		}
		operator := p.previous()
		right := p.comparison()
		expr = Binary{left: expr, operator: operator, right: right}
	}

	return expr
}

func (p *Parser) comparison() Expr {
	expr := p.term()

	for {
		if !p.match([]TokenType{Greater, GreaterEqual, Less, LessEqual}) {
			break
		}
		operator := p.previous()
		right := p.term()
		expr = Binary{left: expr, operator: operator, right: right}
	}

	return expr
}

func (p *Parser) term() Expr {
	expr := p.factor()

	for {
		if !p.match([]TokenType{Minus, Plus}) {
			break
		}

		operator := p.previous()
		right := p.factor()
		expr = Binary{left: expr, operator: operator, right: right}
	}

	return expr
}

func (p *Parser) factor() Expr {
	expr := p.unary()

	for {
		if !p.match([]TokenType{Slash, Star}) {
			break
		}
		operator := p.previous()
		right := p.unary()
		expr = Binary{left: expr, operator: operator, right: right}
	}

	return expr
}

func (p *Parser) match(types []TokenType) bool {
	for _, t := range types {
		if p.check(t) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(t TokenType) bool {
	if p.isAtEnd() {
		return false
	}

	return t == p.peek().Type
}

func (p *Parser) advance() LoxToken {
	if !p.isAtEnd() {
		p.current++
	}
	return p.previous()
}

func (p *Parser) isAtEnd() bool {
	return p.peek().Type == Eof
}

func (p *Parser) peek() LoxToken {
	return p.tokens[p.current]
}

func (p *Parser) previous() LoxToken {
	return p.tokens[p.current-1]
}

func (p *Parser) unary() Expr {
	if p.match([]TokenType{Bang, Minus}) {
		operator := p.previous()
		right := p.unary()
		return Unary{operator: operator, right: right}
	}

	return p.primary()
}

func (p *Parser) primary() Expr {
	if p.match([]TokenType{False}) {
		return Literal{value: false}
	}
	if p.match([]TokenType{True}) {
		return Literal{value: true}
	}
	if p.match([]TokenType{Nil}) {
		return Literal{value: nil}
	}

	if p.match([]TokenType{Number, String}) {
		return Literal{p.previous().Literal} // why previous? that seems unclear
	}

	if p.match([]TokenType{LeftParen}) {
		expr := p.expression()
		p.consume(RightParen, "Expect ')' after expression.")
		return Grouping{expression: expr}
	}

	// throw error(peek(), "Expect exprpssion.")
	// TODO: Implement this in a way that doesn't reach this part.
	panic("Wasn't supposed to reach this far.")
}

func (p *Parser) consume(t TokenType, message string) LoxToken {
	if p.check(t) {
		return p.advance()
	}

	panic(message)
}

func (p *Parser) synchronize() {
	p.advance()

	for {
		if p.current >= int64(len(p.tokens)) {
			return
		}

		if p.previous().Type == Semicolon {
			return
		}

		switch p.peek().Type {
		case Class:
		case Fun:
		case Var:
		case For:
		case If:
		case While:
		case Print:
		case Return:
			return
		}

		p.advance()
	}
}
