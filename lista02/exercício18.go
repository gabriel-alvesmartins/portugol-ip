package main

import (
	"fmt"
)

func main() {
	var (
		d int
		p float64
		t string
	)
	fmt.Println("Qual é o dia da semana?(1,2,3,4,5,6,7 - SENDO 1 DOMINGO E 7 SÁBADO)")
	fmt.Println("Qual o preço normal dos DVD's?")
	fmt.Println("Qual a categoria do DVD?(COMUM {C} OU LANÇAMENTO {L}")
	fmt.Scan(&d, &p, &t)
	switch d {
	case 2, 3, 5:
		if t == "C" {
			p = p * 0.6
		} else {
			p = p * 0.6 * 1.15
		}
	case 4, 6, 7, 1:
		if t == "L" {
			p = p * 1.15
		} else {
		}
	default:
		fmt.Println("Insira um dia válido")
	}
	fmt.Println("O preço final: ", p)
}
