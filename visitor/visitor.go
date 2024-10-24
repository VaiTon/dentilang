package visitor

import (
	"fmt"

	"dentilang/ast"
)

type Visitor interface {
	ast.Visitor
	Result() string
}

type SExprVisitor struct {
	result string
}

func (v *SExprVisitor) Result() string { return v.result }
func (v *SExprVisitor) visitBinaryOperator(op string, left, right ast.Exp) {
	left.Accept(v)
	l := v.result
	right.Accept(v)
	r := v.result

	v.result = fmt.Sprintf("(%s %s %s)", op, l, r)

}
func (v *SExprVisitor) VisitNum(exp ast.NumExp)       { v.result = fmt.Sprintf("%.2f", exp.Value()) }
func (v *SExprVisitor) VisitLIdent(exp ast.LIdentExp) { v.result = ":" + exp.Value() }
func (v *SExprVisitor) VisitRIdent(exp ast.RIdentExp) { v.result = exp.Value() }

func (v *SExprVisitor) VisitSum(exp ast.SumExp) { v.visitBinaryOperator("+", exp.Left(), exp.Right()) }
func (v *SExprVisitor) VisitDiv(exp ast.DivExp) { v.visitBinaryOperator("/", exp.Left(), exp.Right()) }
func (v *SExprVisitor) VisitSub(exp ast.SubExp) { v.visitBinaryOperator("-", exp.Left(), exp.Right()) }
func (v *SExprVisitor) VisitMul(exp ast.MulExp) { v.visitBinaryOperator("*", exp.Left(), exp.Right()) }
func (v *SExprVisitor) VisitPot(exp ast.PotExp) { v.visitBinaryOperator("^", exp.Left(), exp.Right()) }
func (v *SExprVisitor) VisitAssign(exp ast.AssignExp) {
	v.visitBinaryOperator("←", exp.Left(), exp.Right())
}
func (v *SExprVisitor) VisitSequence(exp ast.SequenceExp) {
	exp.Left().Accept(v)
	left := v.result
	exp.Right().Accept(v)
	right := v.result

	v.result = fmt.Sprintf("(%s, %s)", left, right)
}
