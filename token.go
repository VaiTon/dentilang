package main

import (
	"strconv"
)

type TokenType int

const (
	Invalid TokenType = iota
	EOF
	Number
	Plus
	Minus
	Multiply
	Divide
	LeftParen
	RightParen
	Assign
	Semicolon
	Comma
	Identifier
	LValue
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) getNumber() (float64, error) {
	return strconv.ParseFloat(t.Value, 64)
}
