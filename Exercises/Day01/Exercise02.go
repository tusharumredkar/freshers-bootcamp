package main

import "fmt"

/**
Create an expression tree using golang structs. For example string a + b - c on tree would look like the following,
  +
/   \
a   -
    /  \
   b   c


Building the tree can be done manually through code. Once the tree is built, traverse the tree in preorder and postorder format and print the values.


Expectation: Use Recursion, Strings, Pointers, Struct
*/

type Node struct {
	left  *Node
	right *Node
	value string
}

func (node *Node) printTreeInPreOrder() {
	if node == nil {
		return
	}
	fmt.Print(node.value + " ")
	node.left.printTreeInPreOrder()
	node.right.printTreeInPreOrder()
}

func (node *Node) printTreeInPostOrder() {
	if node == nil {
		return
	}
	node.left.printTreeInPostOrder()
	node.right.printTreeInPostOrder()
	fmt.Print(node.value + " ")
}

func (node *Node) printTreeInOrder() {
	if node == nil {
		return
	}
	node.left.printTreeInOrder()
	fmt.Print(node.value + " ")
	node.right.printTreeInOrder()
}

func main() {

	tree_input := []string{"+", "a", "-1", "-1", "-", "b", "-1", "-1", "c", "-1", "-1"}
	index := 0

	node := buildTreePreorder(tree_input, &index)

	fmt.Println("Print in preorder")
	node.printTreeInPreOrder()
	fmt.Println("\n")

	fmt.Println("Print in order")
	node.printTreeInOrder()
	fmt.Println("\n")

	fmt.Println("Print in postorder")
	node.printTreeInPostOrder()
	fmt.Println("")
}

func buildTreePreorder(input []string, index *int) *Node {
	if *index == len(input) || input[*index] == "-1" {
		*index++
		return nil
	}
	node := &Node{value: input[*index]}
	*index++
	node.left = buildTreePreorder(input, index)
	node.right = buildTreePreorder(input, index)
	return node
}
