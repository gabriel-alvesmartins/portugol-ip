package main

import "fmt"

func main() {
	var a, b, pa, pb, md float64

	fmt.Println("Digite duas notas e seus respectivos pesos:")
	fmt.Scan(&a, &b, &pa, &pb)
	md = (a*pa + b*pb) / (pa + pb)
	fmt.Println("A média é igual a: ", md)

}
