package main

import "fmt"

func main() {
	var n int
	var media, soma float64
	fmt.Println("Quantos números você quer na sua média?")
	fmt.Scan(&n)
	soma = 0
	for i := 0; i < n; i++ {
		var number float64
		fmt.Println("Digite o", i+1, "º número")
		fmt.Scan(&number)
		soma = soma + number
	}
	media = soma / float64(n)
	fmt.Println("A média é: ", media)
}
