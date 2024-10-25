package main

import (
	"fmt"
	"testing"

	"dentilang/ast"
	"dentilang/parser"
	visitor2 "dentilang/visitor"
)

func Test_interpret(t *testing.T) {

	exps := []struct {
		expr   string
		result ast.Exp
	}{
		{
			"1 + 2 * 3",
			ast.NewSumExp(
				ast.NewNumExp(1),
				ast.NewMulExp(ast.NewNumExp(2), ast.NewNumExp(3)),
			),
		},
		{
			"(2+3) * 20 / ((3+2)*3)",
			ast.NewDivExp(
				ast.NewMulExp(
					ast.NewSumExp(ast.NewNumExp(2), ast.NewNumExp(3)),
					ast.NewNumExp(20),
				),
				ast.NewMulExp(
					ast.NewSumExp(ast.NewNumExp(3), ast.NewNumExp(2)),
					ast.NewNumExp(3),
				),
			),
		},
	}

	for _, exp := range exps {
		t.Run(exp.expr, func(t *testing.T) {
			sc := parser.NewScanner(exp.expr)
			parser := parser.NewParser(sc)

			ast, err := parser.Parse()
			if err != nil {
				t.Fatal(err)
			}

			if ast != exp.result {
				t.Error("AST failed:", ast, "!=", exp.result)
			}

			v := &visitor2.SExprVisitor{}
			ast.Accept(v)
			sexpr := v.Result()

			fmt.Println(sexpr)
		})
	}
}

func TestYes(t *testing.T) {
	expr := "(1 + 2) * 3"

	sc := parser.NewScanner(expr)
	parser := parser.NewParser(sc)

	ast, err := parser.Parse()
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
