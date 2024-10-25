package svm

import "testing"

func Test_PopPush(t *testing.T) {
	svm := NewStackVirtualMachine()
	if len(svm.stack) != 0 {
		t.Errorf("Expected empty stack, got %v", svm.stack)
	}

	_, err := svm.Pop()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(1)
	if len(svm.stack) != 1 {
		t.Errorf("Expected stack with 1 element, got %v", svm.stack)
	}

	f, err := svm.Pop()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if f != 1 {
		t.Errorf("Expected 1, got %v", f)
	}
}

func Test_Add(t *testing.T) {
	svm := NewStackVirtualMachine()

	err := svm.Add()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(1)
	err = svm.Add()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(2)
	err = svm.Add()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if len(svm.stack) != 1 {
		t.Errorf("Expected stack with 1 element, got %v", svm.stack)
	}

	f := svm.stack[0]
	if f != 3 {
		t.Errorf("Expected 3, got %v", f)
	}
}

func Test_Sub(t *testing.T) {
	svm := NewStackVirtualMachine()

	err := svm.Sub()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(1)
	err = svm.Sub()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(2)
	err = svm.Sub()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if len(svm.stack) != 1 {
		t.Errorf("Expected stack with 1 element, got %v", svm.stack)
	}

	f := svm.stack[0]
	if f != -1 {
		t.Errorf("Expected -1, got %v", f)
	}
}

func Test_Mul(t *testing.T) {
	svm := NewStackVirtualMachine()

	err := svm.Mul()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(1)
	err = svm.Mul()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(2)
	err = svm.Mul()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if len(svm.stack) != 1 {
		t.Errorf("Expected stack with 1 element, got %v", svm.stack)
	}

	f := svm.stack[0]
	if f != 2 {
		t.Errorf("Expected 2, got %v", f)
	}
}

func Test_Div(t *testing.T) {
	svm := NewStackVirtualMachine()

	err := svm.Div()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(1)
	err = svm.Div()
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	svm.Push(2)
	err = svm.Div()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if len(svm.stack) != 1 {
		t.Errorf("Expected stack with 1 element, got %v", svm.stack)
	}

	f := svm.stack[0]
	if f != 0.5 {
		t.Errorf("Expected 0.5, got %v", f)
	}
}

func Test_Stack(t *testing.T) {
	svm := NewStackVirtualMachine()

	s := svm.Stack()
	if len(s) != 0 {
		t.Errorf("Expected empty stack, got %v", s)
	}

	svm.Push(1)
	svm.Push(2)
	svm.Push(3)

	s = svm.Stack()
	if len(s) != 3 {
		t.Errorf("Expected stack with 3 elements, got %v", s)
	}

	if s[0] != 1 || s[1] != 2 || s[2] != 3 {
		t.Errorf("Expected [1 2 3], got %v", s)
	}
}

func Test_Run(t *testing.T) {
	svm := NewStackVirtualMachine()

	err := svm.Run("")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	err = svm.Run("push")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	err = svm.Run("push asd")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	err = svm.Run("asd")
	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	assembler := []string{
		"push 1",
		"push 2",
		"add",
		"push 3",
		"mul",
		"push 4",
		"div",
		"push 1",
		"sub",
	}

	for _, op := range assembler {
		err := svm.Run(op)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	}

	if len(svm.stack) != 1 {
		t.Errorf("Expected stack with 1 element, got %v", svm.stack)
	}

	f := svm.stack[0]
	if f != 1.25 {
		t.Errorf("Expected 1.25, got %v", f)
	}

}
