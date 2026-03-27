package main

import (
	"fmt"
)

func main() {
	var t string
	var conta int
	var pf, cons float64
	fmt.Println("Qual é o seu tipo de consumidor (R,C,I), o número da sua conta e o seu consumo mensal (m³), respectivamente? ")
	fmt.Scan(&t, &conta, &cons)
	switch t {
	case "R":
		pf = 5 + 0.05*cons
	case "C":
		if cons <= 80 {
			pf = 500
		} else {
			pf = 500 + (cons-80)*0.25
		}
	case "I":
		if cons <= 100 {
			pf = 800
		} else {
			pf = 800 + (cons-100)*0.04
		}
	}
	fmt.Println("Sua conta é:", conta)
	fmt.Println("O preço final é:", pf)
}
