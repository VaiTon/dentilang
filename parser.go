package main

import (
	"errors"

	"dentilang/ast"
)

type Parser struct {
	scanner *Scanner
}

func NewParser(scanner *Scanner) *Parser {
	return &Parser{scanner: scanner}
}

func (p *Parser) parseExp() (ast.Exp, error) {
	t1, err := p.parseTerm()
	if err != nil {
		return nil, err
	}

	for (p.scanner.currToken != Token{}) {
		if p.scanner.currToken.Type == Plus {
			p.scanner.nextToken()

			t2, err := p.parseTerm()
			if err != nil {
				return nil, err
			}

			t1 = ast.NewSumExp(t1, t2)
		} else if p.scanner.currToken.Type == Minus {
			p.scanner.nextToken()

			t2, err := p.parseTerm()
			if err != nil {
				return nil, err
			}

			t1 = ast.NewSubExp(t1, t2)
		} else if p.scanner.currToken.Type == Assign {
			p.scanner.nextToken()
			t2, err := p.parseExp()
			if err != nil {
				return nil, err
			}

			return ast.NewAssignExp(t1, t2), nil
		} else {
			return t1, nil
		}
	}

	return t1, nil
}

func (p *Parser) parseTerm() (ast.Exp, error) {
	t1, err := p.parseFactor()
	if err != nil {
		return nil, err // short circuit
	}

	for (p.scanner.currToken != Token{}) {
		if p.scanner.currToken.Type == Multiply {

			p.scanner.nextToken()
			t2, err := p.parseFactor()
			if err != nil {
				return nil, err // short circuit
			}

			t1 = ast.NewMulExp(t1, t2)

		} else if p.scanner.currToken.Type == Divide {
			p.scanner.nextToken()
			t2, err := p.parseFactor()
			if err != nil {
				return nil, err // short circuit
			}

			t1 = ast.NewDivExp(t1, t2)
		} else {
			return t1, nil
		}
	}

	return t1, nil
}

func (p *Parser) parseFactor() (ast.Exp, error) {
	if p.scanner.currToken.Type == LeftParen {
		p.scanner.nextToken()
		innerExp, err := p.parseExp() // self-embedding
		if err != nil {
			return nil, err
		}

		if p.scanner.currToken.Type == RightParen {
			// parentesi irrilevanti
			p.scanner.nextToken()
			return innerExp, nil
		}

		// manca la parentesi chiusa
		return nil, errors.New("missing closing parenthesis")
	} else if p.scanner.currToken.Type == LValue {
		p.scanner.nextToken()
		id := p.scanner.currToken

		p.scanner.nextToken()
		return ast.NewLIdentExp(id.Value), nil
	}

	numb, err := p.scanner.currToken.getNumber()
	if err == nil { // dev’essere un numero
		// non è un fattore, quindi
		p.scanner.nextToken()
		return ast.NewNumExp(numb), nil
	}

	// non è qualcosa di riconosciuto
	return nil, errors.New("not a factor")
}
