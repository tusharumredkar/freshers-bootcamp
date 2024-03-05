package main

import "fmt"

func main() {
	n := 1
	fmt.Println(factorial(n))

	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)
	}
	fmt.Println(fib(7))
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}
