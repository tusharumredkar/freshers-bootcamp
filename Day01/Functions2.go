package main

import "fmt"

func swap(a string, b string) (string, string) {
	return b, a
}

func main() {
	a := "Tushar"
	b := "Umredkar"
	c, d := swap(a, b)
	fmt.Println("initial : ", a, b)
	fmt.Println("final : ", c, d)

}
