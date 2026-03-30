package main

import "fmt"

func main() {
	var (
		n           [3]float64
		soma, media float64
	)
	for i := 0; i < len(n); i++ {
		fmt.Scan(&n[i])
		soma += n[i]
	}
	media = soma / float64(len(n))
	fmt.Println("Média:", media)
	fmt.Println("Números acima da média: ")
	for i := 0; i < len(n); i++ {
		if n[i] > media {
			fmt.Println(n[i])
		}
	}
}
