package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n > 0 {
		fmt.Println("O número é POSITIVO")
	} else if n < 0 {
		fmt.Println("O número é NEGATIVO")
	} else {
		fmt.Println("O número é NULO")
	}
}
