package main

import (
	"fmt"
)

func main() {
	var (
		p, pf float64
		e     int
	)
	fmt.Println("Qual o preço do produto?")
	fmt.Scan(&p)
	fmt.Println("Escolha uma condição de pagamento: \n1- À vista, dinheiro ou cheque \n2- À vista, cartão de crédito \n3- Em 2 vezes \n4- Em 3 vezes")
	fmt.Scan(&e)
	switch e {
	case 1:
		pf = p * 0.9
	case 2:
		pf = p * 0.95
	case 3:
		pf = p
	case 4:
		pf = 1.1 * p
	default:
		fmt.Println("Escolha uma opção válida")
	}
	fmt.Println("O seu preço final é: ", pf)
}
