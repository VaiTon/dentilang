package parser

import (
	"log/slog"
	"regexp"
)

type Scanner struct {
	exp string
	idx int

	currToken Token
	r         *regexp.Regexp
}

var tokenRegex = regexp.MustCompile(`\d+|\+|-|\*|/|\(|\)|=|;|,|:|[a-zA-Z_]\w*`)

func NewScanner(expression string) *Scanner {
	sc := &Scanner{
		exp: expression,
		idx: 0,
		r:   tokenRegex,
	}
	sc.nextToken()
	return sc
}

func (s *Scanner) nextToken() {
	if s.idx >= len(s.exp) {
		s.currToken = Token{Type: EOF}
		return
	}

	m := s.r.FindStringIndex(s.exp[s.idx:])
	if m == nil {
		s.currToken = Token{
			Type:  Invalid,
			Value: s.exp[s.idx:],
		}
		s.idx++
		return
	}

	s.currToken = Token{
		Type:  s.tokenType(s.exp[s.idx+m[0] : s.idx+m[1]]),
		Value: s.exp[s.idx+m[0] : s.idx+m[1]],
	}
	s.idx += m[1]

	slog.Debug("found", "token", s.currToken)
}

var tokenTypeMap = map[string]TokenType{
	"+": Plus,
	"-": Minus,
	"*": Multiply,
	"/": Divide,
	"(": LeftParen,
	")": RightParen,
	"=": Assign,
	";": Semicolon,
	",": Comma,
	":": LValue,
}

var identifierRegex = regexp.MustCompile(`[a-zA-Z_]\w*`)
var numberRegex = regexp.MustCompile(`\d+`)

func (s *Scanner) tokenType(t string) TokenType {
	if tt, ok := tokenTypeMap[t]; ok {
		return tt
	} else if identifierRegex.MatchString(t) {
		return Identifier
	} else if numberRegex.MatchString(t) {
		return Number
	}

	return Invalid
}
