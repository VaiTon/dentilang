# dentilang

An implementation of the compiler studied during the "Linguaggi e Modelli Computazionali" course at the University of Bologna.

## Status

The project is the result of a university course and it is not intended to be used in production. The code is not optimized and it is not guaranteed to be bug-free.

With that said, a list of implemented / missing features is provided below:

### Implemented

- Lexer [scanner.go](./scanner.go)
- Parser [parser.go](./parser.go)
- Abstract Syntax Tree [ast.go](./ast.go)
- Visitors:
  - DOT Visitor [dot.go](./visitor/dot.go)
  - Eval Visitor [eval.go](./visitor/eval.go)
  - Stack Virtual Machine Visitor [svm.go](./visitor/svm.go)

## Getting Started

### Prerequisites

To use the compiler you need to have installed the following software:

- [go](https://golang.org/)
- [git](https://git-scm.com/)

First of all, you need to clone the repository:

```shell
git clone https://github.com/VaiTon/dentilang.git
```

Then you need to install the dependencies and build the project:

```shell
go mod download
go build .
```

### Usage

To run the compiler you need to execute the following command:

```shell
./dentilang [-m <mode>] <expression>
```

Where `<mode>` (default `sexpr`) is one of the following:

- `sexpr`: generates an S-Expression representing the AST and print it to stdout
- `dot`: generates a DOT file representing the AST and print it to stdout
- `eval`: evaluates the expression and print the result to stdout
- `svm`: generates a stack virtual machine code and print it to stdout (see [Stack Virtual Machine](#stack-virtual-machine))

And `<expression>` is the expression to be compiled.

## Stack Virtual Machine

The project includes a simple stack virtual machine that can be used to execute the "compiled" code.

The VM instruction set is the following:

- `PUSH <value>`: push a value on the stack
- `POP`: pop a value from the stack
- `ADD`: pop two values from the stack, add them and push the result
- `SUB`: pop two values from the stack, subtract them and push the result
- `MUL`: pop two values from the stack, multiply them and push the result
- `DIV`: pop two values from the stack, divide them and push the result

I still have doubts about how to implement variable assignment in the stack virtual machine, so it is not implemented yet.

### Running the Stack Virtual Machine

The SVM is located inside the `svm` package and can be executed using the following command:

```shell
go build -o dsvm ./svm
```

and then:

```shell
./dsvm <svm-code>
# or
./dentilang -m svm <expression> | ./dsvm
```

Where `<svm-code>` is the SVM code to be executed, where each instruction is separated by a newline.
