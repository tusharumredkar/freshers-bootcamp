package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func main() {
	a := 1
	b := 2
	c := add(a, b)
	fmt.Println("Sum = ", c)

}
