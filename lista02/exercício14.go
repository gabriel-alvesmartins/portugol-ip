package main

import "fmt"

func main() {
	var a, b, c, d string
	var a1, b1, c1, d1, preco, precof int
	fmt.Println("Qual o valor inicial do carro?")
	fmt.Scan(&preco)
	fmt.Println("Quais opções você quer no carro?")
	fmt.Println("A- Ar Condicionado (S/N)")
	fmt.Scan(&a)
	fmt.Println("B-Pintura Metálica (S/N)")
	fmt.Scan(&b)
	fmt.Println("C- Vidro Elétrico(S/N)")
	fmt.Scan(&c)
	fmt.Println("D- Direção Hidráulica(S/N)")
	fmt.Scan(&d)
	if a == "S" {
		a1 = 1
	} else {
		a1 = 0
	}
	if b == "S" {

		b1 = 1
	} else {

		b1 = 0
	}
	if c == "S" {

		c1 = 1
	} else {

		c1 = 0
	}
	if d == "S" {

		d1 = 1
	} else {

		d1 = 0
	}
	precof = preco + (a1 * 1750) + (800 * b1) + (1200 * c1) + (d1 * 2000)
	fmt.Println("O preço final do carro é:", precof)
}
