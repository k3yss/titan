package main

import (
	"bufio"
	"fmt"
	"os"
)

var hadError bool = false

func main() {
	argsLen := len(os.Args)
	fmt.Println("len", os.Args)
	if argsLen > 2 {
		fmt.Println("Usage: titan [script]")
	} else if argsLen == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	fmt.Println("Path: ", path)
}

func runPrompt() {
	fmt.Println("Running a prompt")

	//currently thinking without the newline character
	scanner := bufio.NewScanner(os.Stdin)
	for {

		fmt.Print(">>")
		scanned := scanner.Scan()
		line := scanner.Text()

		if !scanned {
			break
		}
		run(line)

		hadError = false

	}
}

func run(line string) {
	fmt.Printf("Reading from the run function: %v \n", line)
}
