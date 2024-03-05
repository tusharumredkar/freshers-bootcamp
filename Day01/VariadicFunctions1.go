package main

import "fmt"

func sum(nums ...int) int {

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println("Total: ", total)

	return total
}

func main() {
	sum(1, 2, 5)

}
