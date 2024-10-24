package visitor

import (
	"fmt"
	"strings"

	"dentilang/ast"
)

// SVMVisitor is a visitor that generates "assembly" code
//
// It visits the AST in post-order and generates the assembly code
// for a generic stack-based virtual machine.
type SVMVisitor struct{ result string }

func (v *SVMVisitor) addOp(op ...string) { v.result += strings.Join(op, " ") + "\n" }

func (v *SVMVisitor) visitBinaryOperator(op string, exp ast.OpExp) {
	exp.Left().Accept(v)
	exp.Right().Accept(v)

	v.addOp(op)
}

func (v *SVMVisitor) Result() string { return v.result }

func (v *SVMVisitor) VisitLIdent(exp ast.LIdentExp) { v.addOp("DECLARE", exp.Value()) }
func (v *SVMVisitor) VisitRIdent(exp ast.RIdentExp) { v.addOp("LOAD", exp.Value()) }
func (v *SVMVisitor) VisitNum(exp ast.NumExp)       { v.addOp("PUSH", fmt.Sprintf("%f", exp.Value())) }

func (v *SVMVisitor) VisitSum(exp ast.SumExp)     { v.visitBinaryOperator("ADD", exp) }
func (v *SVMVisitor) VisitDiv(exp ast.DivExp)     { v.visitBinaryOperator("DIV", exp) }
func (v *SVMVisitor) VisitSub(exp ast.SubExp)     { v.visitBinaryOperator("SUB", exp) }
func (v *SVMVisitor) VisitMul(exp ast.MulExp)     { v.visitBinaryOperator("MUL", exp) }
func (v *SVMVisitor) VisitPot(exp ast.PotExp)     { v.visitBinaryOperator("POT", exp) }
func (v *SVMVisitor) VisitAssign(e ast.AssignExp) { v.visitBinaryOperator("ASSIGN", e) }

func (v *SVMVisitor) VisitSequence(e ast.SequenceExp) {
	e.Left().Accept(v)
	e.Right().Accept(v)
}
