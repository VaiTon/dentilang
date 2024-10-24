package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/fatih/color"

	visitor2 "dentilang/visitor"
)

type LevelColorHandler struct {
	slog.Handler
	debug bool
}

func (h *LevelColorHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl != slog.LevelDebug || h.debug
}

func (h *LevelColorHandler) Handle(_ context.Context, r slog.Record) error {

	var colorFunc func(format string, a ...interface{}) string
	switch r.Level {
	case slog.LevelDebug:
		colorFunc = color.MagentaString
	case slog.LevelInfo:
		colorFunc = color.GreenString
	case slog.LevelWarn:
		colorFunc = color.YellowString
	case slog.LevelError:
		colorFunc = color.RedString
	default:
		colorFunc = color.WhiteString
	}

	msg := r.Message

	lvl := r.Level.String()
	lvl = colorFunc("%-5s:", strings.ToLower(lvl))

	grayFunc := color.New(color.FgHiBlack).SprintfFunc()

	msg = fmt.Sprintf("%s %s", lvl, msg)
	// add args
	r.Attrs(func(a slog.Attr) bool {
		msg += fmt.Sprintf(" %s=%v", grayFunc(a.Key), a.Value)
		return true
	})

	fmt.Println(msg)
	return nil
}

var (
	verbose = flag.Bool("v", false, "Enable verbose logging")
	visitor = flag.String("m", "", "Select visitor")
)

func main() {
	flag.Parse()
	slog.SetDefault(slog.New(&LevelColorHandler{debug: *verbose}))

	var v visitor2.Visitor
	if *visitor == "" || *visitor == "sexpr" {
		v = &visitor2.SExprVisitor{}
	} else if *visitor == "eval" {
		v = &visitor2.EvalVisitor{}
	} else if *visitor == "svm" {
		v = &visitor2.SVMVisitor{}
	} else if *visitor == "dot" {
		v = &visitor2.DotVisitor{}
	} else {
		slog.Error("unknown visitor: ", "visitor", *visitor)
		os.Exit(1)
	}

	if flag.NArg() != 1 {
		slog.Error(fmt.Sprintf("usage: %s [options] <expression>", os.Args[0]))
		os.Exit(1)
	}

	expr := flag.Arg(0)

	sc := NewScanner(expr)
	parser := NewParser(sc)

	ast, err := parser.parseExp()
	if err != nil {
		slog.Error(fmt.Sprintf("syntax error: %v", err))
		os.Exit(1)
	}

	ast.Accept(v)

	fmt.Println(v.Result())
}
