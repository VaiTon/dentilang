package main

import (
	"fmt"
	"testing"

	"dentilang/ast"
	visitor2 "dentilang/visitor"
)

func Test_interpret(t *testing.T) {

	exps := []struct {
		expr   string
		result ast.Exp
	}{
		{
			"1 + 2 * 3",
			ast.SumExp{ast.NumExp{1}, ast.MulExp{ast.NumExp{2}, ast.NumExp{3}}},
		},
		{
			"(2+3) * 20 / ((3+2)*3)",
			ast.DivExp{
				ast.MulExp{ast.SumExp{ast.NumExp{2}, ast.NumExp{3}}, ast.NumExp{20}},
				ast.MulExp{ast.SumExp{ast.NumExp{3}, ast.NumExp{2}}, ast.NumExp{3}},
			},
		},
	}

	for _, exp := range exps {
		t.Run(exp.expr, func(t *testing.T) {
			sc := NewScanner(exp.expr)
			parser := NewParser(sc)

			ast, err := parser.parseExp()
			if err != nil {
				t.Fatal(err)
			}

			if ast != exp.result {
				t.Error("AST failed:", ast, "!=", exp.result)
			}

			v := &visitor2.SExprVisitor{}
			ast.Accept(v)
			sexpr := v.result

			fmt.Println(sexpr)
		})
	}
}

func TestYes(t *testing.T) {
	expr := "(1 + 2) * 3"

	sc := NewScanner(expr)
	parser := NewParser(sc)

	ast, err := parser.parseExp()
	if err != nil {
		t.Fatal(err)
	}

	var v visitor2.Visitor

	v = &visitor2.SExprVisitor{}
	ast.Accept(v)

	fmt.Println("SExpr\t->", v.Result())

	v = &visitor2.SVMVisitor{}
	ast.Accept(v)

	fmt.Println("SVM\t->", v.Result())

	v = &visitor2.EvalVisitor{}
	ast.Accept(v)

	fmt.Println("Eval\t->", v.Result())

}
