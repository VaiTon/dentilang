package visitor

import (
	"fmt"
	"math"

	"dentilang/ast"
)

type EvalVisitor struct {
	result float64

	memory map[string]float64
}

func (v *EvalVisitor) Result() string { return fmt.Sprintf("%f", v.result) }

func (v *EvalVisitor) VisitNum(e ast.NumExp) {
	v.result = e.Value()
}

func (v *EvalVisitor) VisitSum(e ast.SumExp) {
	e.Left().Accept(v)
	left := v.result
	e.Right().Accept(v)
	right := v.result

	v.result = left + right
}

func (v *EvalVisitor) VisitDiv(e ast.DivExp) {
	e.Left().Accept(v)
	left := v.result
	e.Right().Accept(v)
	right := v.result

	v.result = left / right
}

func (v *EvalVisitor) VisitSub(e ast.SubExp) {
	e.Left().Accept(v)
	left := v.result
	e.Right().Accept(v)
	right := v.result

	v.result = left - right
}

func (v *EvalVisitor) VisitMul(e ast.MulExp) {
	e.Left().Accept(v)
	left := v.result
	e.Right().Accept(v)
	right := v.result

	v.result = left * right
}

func (v *EvalVisitor) VisitPot(e ast.PotExp) {
	e.Left().Accept(v)
	left := v.result
	e.Right().Accept(v)
	right := v.result

	v.result = math.Pow(left, right)
}
func (v *EvalVisitor) VisitLIdent(e ast.LIdentExp) { /* nothing to do */ }
func (v *EvalVisitor) VisitRIdent(e ast.RIdentExp) { /* nothing to do */ }
func (v *EvalVisitor) VisitAssign(e ast.AssignExp) { panic("implement me") }
func (v *EvalVisitor) VisitSequence(e ast.SequenceExp) {
	e.Left().Accept(v)
	e.Right().Accept(v)
}
