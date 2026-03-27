package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, delta, r, r1, r2 float64
	fmt.Scan(&a, &b, &c)
	if a == 0 {
		r = -c / b
		fmt.Println("O resultado da sua equação de primeiro grau é:", r)
		return
	}
	delta = b*b - 4*a*c
	if delta < 0 == false {
		if delta == 0 {
			fmt.Println("Raiz Única")
		} else {
			fmt.Println("Raízes Distintas")
		}
		r1 = (-b + math.Sqrt(delta)) / (2 * a)
		r2 = (-b - math.Sqrt(delta)) / (2 * a)
		if r1 == r2 {
			fmt.Println("O valor da raíz é :", r1)
		} else {
			fmt.Printf("Os valores das raízes são: %f e %f", r1, r2)
		}
	} else {
		fmt.Println("Raízes Imaginárias, não conseguimos calcular no momento.")
	}
}
