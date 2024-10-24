package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "usage: %s <file>\n", os.Args[0])
		return
	}

	file, err := os.Open(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
	}()

	svm := NewStackVirtualMachine()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}

		fmt.Println("OP:", scanner.Text())
		err = svm.Run(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		fmt.Println("Stack:", svm.stack)
	}

	fmt.Println("Execution finished.")
}

type StackVirtualMachine struct {
	stack []float64
}

func NewStackVirtualMachine() *StackVirtualMachine {
	return &StackVirtualMachine{}
}

func (s *StackVirtualMachine) Push(f float64) {
	s.stack = append(s.stack, f)
}

func (s *StackVirtualMachine) Pop() (float64, error) {
	if len(s.stack) == 0 {
		return 0, fmt.Errorf("stack is empty")
	}

	f := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return f, nil
}

func (s *StackVirtualMachine) Add() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b, err := s.Pop()
	if err != nil {
		return err
	}

	a, err := s.Pop()
	if err != nil {
		return err
	}

	s.Push(a + b)
	return nil
}

func (s *StackVirtualMachine) Sub() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b, err := s.Pop()
	if err != nil {
		return err
	}

	a, err := s.Pop()
	if err != nil {
		return err
	}

	s.Push(a - b)
	return nil
}

func (s *StackVirtualMachine) Mul() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b, err := s.Pop()
	if err != nil {
		return err
	}

	a, err := s.Pop()
	if err != nil {
		return err
	}

	s.Push(a * b)
	return nil
}

func (s *StackVirtualMachine) Div() error {
	if len(s.stack) < 2 {
		return fmt.Errorf("stack is empty")
	}

	b, err := s.Pop()
	if err != nil {
		return err
	}

	a, err := s.Pop()
	if err != nil {
		return err
	}

	s.Push(a / b)
	return nil
}

func (s *StackVirtualMachine) Stack() []float64 {
	return s.stack
}

func (s *StackVirtualMachine) Run(op string) error {
	arr := strings.Split(op, " ")
	if len(arr) < 1 {
		return fmt.Errorf("no operation")
	}

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
