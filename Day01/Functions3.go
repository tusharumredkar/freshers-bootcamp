package main

import "fmt"

func split(sum int) (a int, b int) {
	a = sum % 2
	b = sum / 2
	return
}

func main() {

	fmt.Println(split(151))
}
