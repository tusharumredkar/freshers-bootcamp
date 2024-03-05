package main

import (
	"fmt"
)

func main() {
	i := 1
	switch k := i % 2; k {
	case 0:
		fmt.Println("Even")
	case 1:
		fmt.Println("Odd")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("Non integer")
	}
}
