package ast

import (
	"testing"
)

func TestNumExp(t *testing.T) {
	num := NewNumExp(42)
	if num.Value() != 42 {
		t.Errorf("Expected 42, got %v", num.Value())
	}
}

func TestLIdentExp(t *testing.T) {
	ident := NewLIdentExp("foo")
	if ident.Value() != "foo" {
		t.Errorf("Expected foo, got %v", ident.Value())
	}
}

func TestRIdentExp(t *testing.T) {
	ident := NewRIdentExp("foo")
	if ident.Value() != "foo" {
		t.Errorf("Expected foo, got %v", ident.Value())
	}
}

func TestSumExp(t *testing.T) {
	left := NewNumExp(1)
	right := NewNumExp(2)
	sum := NewSumExp(left, right)

	if sum.Left() != left {
		t.Errorf("Expected %v, got %v", left, sum.Left())
	}

	if sum.Right() != right {
		t.Errorf("Expected %v, got %v", right, sum.Right())
	}
}

func TestMulExp(t *testing.T) {
	left := NewNumExp(1)
	right := NewNumExp(2)
	mul := NewMulExp(left, right)

	if mul.Left() != left {
		t.Errorf("Expected %v, got %v", left, mul.Left())
	}

	if mul.Right() != right {
		t.Errorf("Expected %v, got %v", right, mul.Right())
	}
}

func TestDivExp(t *testing.T) {
	left := NewNumExp(1)
	right := NewNumExp(2)
	div := NewDivExp(left, right)

	if div.Left() != left {
		t.Errorf("Expected %v, got %v", left, div.Left())
	}

	if div.Right() != right {
		t.Errorf("Expected %v, got %v", right, div.Right())
	}
}

func TestSubExp(t *testing.T) {
	left := NewNumExp(1)
	right := NewNumExp(2)
	sub := NewSubExp(left, right)

	if sub.Left() != left {
		t.Errorf("Expected %v, got %v", left, sub.Left())
	}

	if sub.Right() != right {
		t.Errorf("Expected %v, got %v", right, sub.Right())
	}
}

func TestPotExp(t *testing.T) {
	left := NewNumExp(1)
	right := NewNumExp(2)
	pot := NewPotExp(left, right)

	if pot.Left() != left {
		t.Errorf("Expected %v, got %v", left, pot.Left())
	}

	if pot.Right() != right {
		t.Errorf("Expected %v, got %v", right, pot.Right())
	}
}

type MockVisitor struct {
	calls map[string]int
}

func (v *MockVisitor) VisitNum(NumExp)           { v.calls["num"]++ }
func (v *MockVisitor) VisitLIdent(LIdentExp)     { v.calls["lident"]++ }
func (v *MockVisitor) VisitRIdent(RIdentExp)     { v.calls["rident"]++ }
func (v *MockVisitor) VisitSum(SumExp)           { v.calls["sum"]++ }
func (v *MockVisitor) VisitDiv(DivExp)           { v.calls["div"]++ }
func (v *MockVisitor) VisitSub(SubExp)           { v.calls["sub"]++ }
func (v *MockVisitor) VisitMul(MulExp)           { v.calls["mul"]++ }
func (v *MockVisitor) VisitPot(PotExp)           { v.calls["pot"]++ }
func (v *MockVisitor) VisitAssign(AssignExp)     { v.calls["assign"]++ }
func (v *MockVisitor) VisitSequence(SequenceExp) { v.calls["sequence"]++ }

func TestVisitor(t *testing.T) {
	v := &MockVisitor{calls: make(map[string]int)}
	num := NewNumExp(42)
	num.Accept(v)
	if v.calls["num"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["num"])
	}

	lident := NewLIdentExp("foo")
	lident.Accept(v)
	if v.calls["lident"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["lident"])
	}

	rident := NewRIdentExp("foo")
	rident.Accept(v)
	if v.calls["rident"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["rident"])
	}

	left := NewNumExp(1)
	right := NewNumExp(2)
	sum := NewSumExp(left, right)
	sum.Accept(v)
	if v.calls["sum"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["sum"])
	}

	div := NewDivExp(left, right)
	div.Accept(v)
	if v.calls["div"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["div"])
	}

	sub := NewSubExp(left, right)
	sub.Accept(v)
	if v.calls["sub"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["sub"])
	}

	mul := NewMulExp(left, right)
	mul.Accept(v)
	if v.calls["mul"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["mul"])
	}

	pot := NewPotExp(left, right)
	pot.Accept(v)
	if v.calls["pot"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["pot"])
	}

	sequence := NewSequenceExp(left, right)
	sequence.Accept(v)
	if v.calls["sequence"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["sequence"])
	}

	assign := NewAssignExp(NewLIdentExp("foo"), NewNumExp(42))
	assign.Accept(v)
	if v.calls["assign"] != 1 {
		t.Errorf("Expected 1, got %v", v.calls["assign"])
	}

}


func TestSequenceExp(t *testing.T) {
	left := NewNumExp(1)
	right := NewNumExp(2)
	sequence := NewSequenceExp(left, right)

	if sequence.Left() != left {
		t.Errorf("Expected %v, got %v", left, sequence.Left())
	}

	if sequence.Right() != right {
		t.Errorf("Expected %v, got %v", right, sequence.Right())
	}
}

func TestAssignExp(t *testing.T) {
	left := NewLIdentExp("foo")
	right := NewNumExp(2)
	assign := NewAssignExp(left, right)

	if assign.Left() != left {
		t.Errorf("Expected %v, got %v", left, assign.Left())
	}

	if assign.Right() != right {
		t.Errorf("Expected %v, got %v", right, assign.Right())
	}
}
