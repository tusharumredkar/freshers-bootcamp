package main

import "fmt"

func main() {
	sum := 1

	if sum = sum * sum; sum < 100 {
		fmt.Println("sum = ", sum)
	} else {
		fmt.Println("Limit exceeded")
	}

}
