package main

import "fmt"

func main() {
	var soma, a1, r float64
	var n int

	fmt.Println("Digite, respectivamente, o valor do primeiro termo, a razão e o número de elementos da progressão:")
	fmt.Scan(&a1, &r, &n)

	soma = 0
	termo := a1
	for i := 0; i < n; i++ {
		soma += termo
		termo += r
	}
	fmt.Println(soma)
}
