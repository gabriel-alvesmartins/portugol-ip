package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	if a%b == 0 {
		fmt.Println("Os números são divisíveis")
	} else {
		fmt.Println("Os números não são divisíveis")
	}
}
