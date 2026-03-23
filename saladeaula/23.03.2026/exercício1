package main

import "fmt"

func main() {
	var a, b, c float64

	fmt.Println("Escreva os valores do lado A,B e C, respectivamente")
	fmt.Scan(&a, &b, &c)
	if (a+b < c) || (a+c < b) || (b+c < a) {
		fmt.Println("Esses lados não constituem um triângulo")
	} else if a != b && a != c && b != c {
		fmt.Println("Esse triângulo é Escaleno.")
	} else if a == b && a == c && b == c {
		fmt.Println("Esse triângulo é Equilátero")
	} else {
		fmt.Println("Esse triângulo é Isosceles")
	}
}
