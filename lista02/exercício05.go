package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if n%5 == 0 {
		fmt.Printf("O número %d é divisível por 5\n", n)
	} else {
		fmt.Printf("O número %d não é divisível por 5\n", n)
	}
}
