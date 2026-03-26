package main

import (
	"fmt"
)

func main() {
	var a, b, c, maior, inter, menor int
	fmt.Scan(&a, &b, &c)
	if a != b && a != c && b != c {
		if a > b && a > c && b > c { // a>b>c
			menor = c
			inter = b
			maior = a

		} else if a > c && a > b && c > b { // a>c>b
			menor = b
			inter = c
			maior = a
		} else if b > a && a > c && b > c { // b>a>c
			menor = c
			inter = a
			maior = b
		} else if b > a && b > c && c > a { // b>c>a
			menor = a
			inter = c
			maior = b
		} else if c > a && c > b && a > b { // c>a>b
			menor = b
			inter = a
			maior = c
		} else if c > b && c > a && b > a { // c>b>a
			menor = a
			inter = b
			maior = c
		}
		fmt.Println(menor, inter, maior)
	} else {
		fmt.Println("Você colocou números iguais.")
	}
}
