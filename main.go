package main

import "fmt"

type Edge struct {
	From  Node
	To    Node
	Value string
}

type Node struct {
	accepted bool
	next     []Edge
}

func createMachine(expression string) {

	for _, value := range expression {
		fmt.Println(value)

	}

}

func main() {
	fmt.Println("vim-go")

	createMachine("abcd")

	//parseRangeQualifier(regex_expression)
}
