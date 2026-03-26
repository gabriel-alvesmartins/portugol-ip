package main

import "fmt"

func main() {
	var x, r, v int
	fmt.Println("Para qual região você quer ir?")
	fmt.Println("1- Região Norte; \n2- Região Nordeste; \n3- Região Centro-Oeste; \n4- Região Sul;")
	fmt.Scan(&x)
	fmt.Println("Sua passagem inclui retorno? \n1-Sim!\n2-Não!")
	fmt.Scan(&r)
	switch x {
	case 1:
		if r == 1 {
			v = 900
		} else {
			v = 500
		}
	case 2:
		if r == 1 {
			v = 650
		} else {
			v = 350
		}
	case 3:
		if r == 1 {
			v = 600
		} else {
			v = 350
		}
	case 4:
		if r == 1 {
			v = 550
		} else {
			v = 300
		}
	default:
		fmt.Println("Insira um valor válido")
	}
	fmt.Println("O valor da sua passagem é R$ ", v)
}
