// Package visitor provides the interface for the visitor pattern.
//
// The visitor pattern is a way to separate an algorithm from an object structure on which it operates.
// It uses the double dispatch mechanism to call the appropriate method on the visitor based on the type of the object.
// This allows adding new operations to the object structure without modifying the objects themselves.
//
// The visitor pattern is used in the AST package to implement different operations on the AST nodes.
// Each operation is implemented as a visitor that implements the Visitor interface.
//
// # Visitors
// At the moment, there are 4 visitors implemented in the visitor package:
// - DotVisitor: generates a DOT representation of the AST
// - SExprVisitor: generates an S-expression representation of the AST
// - EvalVisitor: evaluates the AST and returns the result
// - StackVirtualMachineVisitor: generates code for a stack-based virtual machine
package visitor

import "dentilang/ast"

type Visitor interface {
	ast.Visitor
	Result() string
}
