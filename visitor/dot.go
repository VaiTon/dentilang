package visitor

import (
	"fmt"

	"dentilang/ast"
)

type DotVisitor struct {
	result string
	id     int
}

func (v *DotVisitor) addStatement(s string) {
	v.result += "\n" + s + ";"
}
func (v *DotVisitor) createElement(label string) {
	v.id++
	v.addStatement(fmt.Sprintf("E%d [label=\"%s\"]", v.id, label))
}
func (v *DotVisitor) addBinaryNode(label string, exp ast.OpExp) {
	exp.Left().Accept(v)
	leftId := v.id
	exp.Right().Accept(v)
	rightId := v.id

	v.id++
	v.addStatement(fmt.Sprintf("E%d [label=\"%s\"]", v.id, label))
	v.addStatement(fmt.Sprintf("E%d -> E%d", v.id, leftId))
	v.addStatement(fmt.Sprintf("E%d -> E%d", v.id, rightId))
}

func (v *DotVisitor) Result() string { return fmt.Sprintf("digraph G {%s\n}", v.result) }

func (v *DotVisitor) VisitSum(e ast.SumExp)           { v.addBinaryNode("+", e) }
func (v *DotVisitor) VisitDiv(e ast.DivExp)           { v.addBinaryNode("/", e) }
func (v *DotVisitor) VisitSub(e ast.SubExp)           { v.addBinaryNode("-", e) }
func (v *DotVisitor) VisitMul(e ast.MulExp)           { v.addBinaryNode("*", e) }
func (v *DotVisitor) VisitPot(e ast.PotExp)           { v.addBinaryNode("^", e) }
func (v *DotVisitor) VisitSequence(e ast.SequenceExp) { v.addBinaryNode(",", e) }
func (v *DotVisitor) VisitAssign(e ast.AssignExp)     { v.addBinaryNode("←", e) }

func (v *DotVisitor) VisitNum(e ast.NumExp)       { v.createElement(fmt.Sprintf("%f", e.Value())) }
func (v *DotVisitor) VisitLIdent(e ast.LIdentExp) { v.createElement(e.Value()) }
func (v *DotVisitor) VisitRIdent(e ast.RIdentExp) { v.createElement(":" + e.Value()) }
