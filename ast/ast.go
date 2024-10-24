package ast

type Visitor interface {
	VisitNum(NumExp)
	VisitSum(SumExp)
	VisitDiv(DivExp)
	VisitSub(SubExp)
	VisitMul(MulExp)
	VisitPot(PotExp)
	VisitLIdent(LIdentExp)
	VisitRIdent(RIdentExp)
	VisitAssign(AssignExp)
	VisitSequence(SequenceExp)
}

type Exp interface {
	Accept(Visitor)
}

type OpExp interface {
	Exp
	Left() Exp
	Right() Exp
}

type LIdentExp struct{ ident string }

func NewLIdentExp(ident string) LIdentExp { return LIdentExp{ident} }
func (l LIdentExp) Value() string         { return l.ident }
func (l LIdentExp) Accept(v Visitor)      { v.VisitLIdent(l) }

type RIdentExp struct{ ident string }

func NewRIdentExp(ident string) RIdentExp { return RIdentExp{ident} }
func (r RIdentExp) Value() string         { return r.ident }
func (r RIdentExp) Accept(v Visitor)      { v.VisitRIdent(r) }

type NumExp struct{ num float64 }

func NewNumExp(num float64) NumExp { return NumExp{num} }
func (n NumExp) Value() float64    { return n.num }
func (n NumExp) Accept(v Visitor)  { v.VisitNum(n) }

type SumExp struct{ left, right Exp }

func NewSumExp(left, right Exp) SumExp { return SumExp{left, right} }
func (e SumExp) Left() Exp             { return e.left }
func (e SumExp) Right() Exp            { return e.right }
func (e SumExp) Accept(v Visitor)      { v.VisitSum(e) }

type MulExp struct{ left, right Exp }

func NewMulExp(left, right Exp) MulExp { return MulExp{left, right} }
func (e MulExp) Left() Exp             { return e.left }
func (e MulExp) Right() Exp            { return e.right }
func (e MulExp) Accept(v Visitor)      { v.VisitMul(e) }

type DivExp struct{ left, right Exp }

func NewDivExp(left, right Exp) DivExp { return DivExp{left, right} }
func (e DivExp) Left() Exp             { return e.left }
func (e DivExp) Right() Exp            { return e.right }
func (e DivExp) Accept(v Visitor)      { v.VisitDiv(e) }

type SubExp struct{ left, right Exp }

func NewSubExp(left, right Exp) SubExp { return SubExp{left, right} }
func (s SubExp) Left() Exp             { return s.left }
func (s SubExp) Right() Exp            { return s.right }
func (s SubExp) Accept(v Visitor)      { v.VisitSub(s) }

type PotExp struct{ left, right Exp }

func NewPotExp(left, right Exp) PotExp { return PotExp{left, right} }
func (p PotExp) Left() Exp             { return p.left }
func (p PotExp) Right() Exp            { return p.right }
func (p PotExp) Accept(v Visitor)      { v.VisitPot(p) }

type AssignExp struct{ left, right Exp }

func NewAssignExp(left, right Exp) AssignExp { return AssignExp{left, right} }
func (a AssignExp) Left() Exp                { return a.left }
func (a AssignExp) Right() Exp               { return a.right }
func (a AssignExp) Accept(v Visitor)         { v.VisitAssign(a) }

type SequenceExp struct{ left, right Exp }

func NewSequenceExp(left, right Exp) SequenceExp { return SequenceExp{left, right} }
func (a SequenceExp) Left() Exp                  { return a.left }
func (a SequenceExp) Right() Exp                 { return a.right }
func (a SequenceExp) Accept(v Visitor)           { v.VisitSequence(a) }
