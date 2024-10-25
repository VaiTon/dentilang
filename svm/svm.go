package svm

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type StackVirtualMachine struct {
	stack []float64
}

func NewStackVirtualMachine() *StackVirtualMachine {
	return &StackVirtualMachine{}
}

// Push pushes a float64 onto the stack
func (s *StackVirtualMachine) Push(f float64) {
	s.stack = append(s.stack, f)
}

// Pop pops the top element of the stack.
//
// If the stack is empty, an error is returned.
func (s *StackVirtualMachine) Pop() (float64, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	return s.unsafePop(), nil
}

// unsafePop pops the top element of the stack without checking if the stack is empty
//
// This function is unsafe and should only be called when the stack is known to be non-empty.
// Instead, use Pop() to safely pop the top element of the stack.
func (s *StackVirtualMachine) unsafePop() float64 {
	f := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return f
}

// Add adds the top two elements of the stack
func (s *StackVirtualMachine) Add() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b := s.unsafePop()
	a := s.unsafePop()

	s.Push(a + b)
	return nil
}

// Sub subtracts the top element of the stack from the second top element
func (s *StackVirtualMachine) Sub() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b := s.unsafePop()
	a := s.unsafePop()

	s.Push(a - b)
	return nil
}

// Mul multiplies the top two elements of the stack
func (s *StackVirtualMachine) Mul() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b := s.unsafePop()
	a := s.unsafePop()

	s.Push(a * b)
	return nil
}

// Div divides the top element of the stack by the second top element
func (s *StackVirtualMachine) Div() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b := s.unsafePop()
	a := s.unsafePop()

	s.Push(a / b)
	return nil
}

// Stack returns a copy of the stack
func (s *StackVirtualMachine) Stack() []float64 {
	return slices.Clone(s.stack)
}

func (s *StackVirtualMachine) Run(op string) error {
	arr := strings.Split(op, " ")

	opcode := arr[0]
	opcode = strings.ToLower(opcode)

	switch opcode {
	case "push":
		if len(arr) < 2 {
			return fmt.Errorf("no operand")
		}
		f, err := strconv.ParseFloat(arr[1], 64)
		if err != nil {
			return err
		}
		s.Push(f)
	case "add":
		return s.Add()
	case "sub":
		return s.Sub()
	case "mul":
		return s.Mul()
	case "div":
		return s.Div()
	default:
		return fmt.Errorf("unknown opcode: %s", opcode)
	}

	return nil
}
