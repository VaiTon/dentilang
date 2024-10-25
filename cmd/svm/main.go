package main

import (
	"bufio"
	"dentilang/svm"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	var scanner *bufio.Scanner

	if len(args) == 0 {
		scanner = bufio.NewScanner(os.Stdin)
	} else {

		file, err := os.Open(args[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}

		scanner = bufio.NewScanner(file)

		defer func() {
			err = file.Close()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error: %v\n", err)
			}
		}()
	}

	svm := svm.NewStackVirtualMachine()

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}

		fmt.Println("OP:", scanner.Text())
		err := svm.Run(scanner.Text())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
		}
		fmt.Println("Stack:", svm.Stack())
	}

	fmt.Println("Execution finished.")
}
