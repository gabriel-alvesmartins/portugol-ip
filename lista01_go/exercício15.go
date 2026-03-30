package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Digite um valor N, tal que N seja maior que 5 e menor que 2000:")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			fmt.Printf("%d^2 = %d\n", i, i*i)
		}
	}
}
