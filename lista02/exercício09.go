package main

import "fmt"

func main() {
	var v, r float64
	fmt.Scan(&v)
	if v > 0 {
		if v < 10 {
			r = v * 1.7
		} else if v >= 10 && v < 30 {
			r = v * 1.5
		} else if v >= 30 && v < 50 {
			r = v * 1.4
		} else {
			r = v * 1.3
		}
		fmt.Println("O valor em reais é: R$", r)
	} else {
		fmt.Println("Valor INVÁLIDO")
	}
}
