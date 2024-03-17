package main

import (
	"strconv"
	"unicode"
)

type LoxScanner struct {
	source  string
	tokens  []LoxToken
	start   int
	current int
	line    int
}

func NewLoxScanner(source string) *LoxScanner {
	return &LoxScanner{source, []LoxToken{}, 0, 0, 1}
}

func (s *LoxScanner) ScanTokens() []LoxToken {
	tokens := make([]LoxToken, 0)

	for {
		if s.current >= len(s.source) {
			break
		}

		s.start = s.current
		s.scanToken()
	}

	tokens = append(tokens, LoxToken{Eof, "", nil, s.line})

	return tokens
}

// Test if the following character is the expected one.
// Only consume character if so.
func (s *LoxScanner) match(expected byte) bool {
	if s.current >= len(s.source) {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++

	return true
}

func (s *LoxScanner) scanToken() {
	c := s.advance()

	switch c {
	case '(':
		s.addToken(LeftParen, nil)

		return
	case ')':
		s.addToken(RightParen, nil)

		return
	case '{':
		s.addToken(LeftBrace, nil)

		return
	case '}':
		s.addToken(RightBrace, nil)

		return
	case ',':
		s.addToken(Comma, nil)

		return
	case '.':
		s.addToken(Dot, nil)

		return
	case '-':
		s.addToken(Minus, nil)

		return
	case '+':
		s.addToken(Plus, nil)

		return
	case ';':
		s.addToken(Semicolon, nil)

		return
	case '*':
		s.addToken(Star, nil)

		return

	case '!':
		if s.match('=') {
			s.addToken(BangEqual, nil)
		} else {
			s.addToken(Bang, nil)
		}

		return

	case '=':
		if s.match('=') {
			s.addToken(EqualEqual, nil)
		} else {
			s.addToken(Equal, nil)
		}

		return

	case '<':
		if s.match('=') {
			s.addToken(LessEqual, nil)
		} else {
			s.addToken(Equal, nil)
		}

		return

	case '>':
		if s.match('=') {
			s.addToken(GreaterEqual, nil)
		} else {
			s.addToken(Greater, nil)
		}

		return

	case '/':
		if s.match('/') {
			for {
				if s.peek() == '\n' || s.current >= len(s.source) {
					break
				}

				s.advance()
			}
		}

		return

	case ' ':
	case '\r':
	case '\t':

		return

	case '\n':
		s.line++

		return

	case '"':
		s.scanString()

		return

	case 'o':
		if s.match('r') {
			s.addToken(Or, nil)
		}

		return

	default:
		if isDigit(c) {
			s.scanNumber()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			reportError(int64(s.line), "Unexpected character.")
		}

		return
	}
}

func isAlpha(b byte) bool {
	return b >= 'a' && b <= 'z' ||
		b >= 'A' && b <= 'Z' ||
		b == '_'
}

func isAlphaNumeric(b byte) bool {
	return isAlpha(b) || isDigit(b)
}

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

func (s *LoxScanner) advance() byte {
	r := s.source[s.current]
	s.current++

	return r
}

func (s *LoxScanner) peek() byte {
	if s.current >= len(s.source) {
		return 0
	}

	return s.source[s.current]
}

func (s *LoxScanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return 0
	}

	return s.source[s.current+1]
}

func (s *LoxScanner) addToken(t TokenType, literal interface{}) {
	text := s.source[s.start:s.current]

	s.tokens = append(s.tokens, LoxToken{t, text, literal, s.line})
}

func (s *LoxScanner) scanString() {
	for {
		if s.peek() != '"' && s.current >= len(s.source) {
			break
		}

		if s.peek() == '\n' {
			s.line++
		}

		s.advance()
	}

	if s.current >= len(s.source) {
		reportError(int64(s.line), "Unterminated string.")

		return
	}

	s.advance()

	s.addToken(String, s.source[s.start+1:s.current-1])
}

func (s *LoxScanner) scanNumber() {
	for {
		if !unicode.IsDigit(rune(s.peek())) {
			break
		}

		s.advance()
	}

	if s.peek() == '.' && unicode.IsDigit(rune(s.peekNext())) {
		s.advance()

		for {
			if !unicode.IsDigit(rune(s.peek())) {
				break
			}

			s.advance()
		}
	}

	f, err := strconv.ParseFloat(s.source[s.start:s.current], 64)
	if err != nil {
		panic(err) // TODO: handle this in a graceful way
	}

	s.addToken(Number, f)
}

func (s *LoxScanner) identifier() {
	for {
		if !isAlphaNumeric(s.peek()) {
			break
		}

		s.advance()
	}

	s.addToken(Identifier, nil)
}
