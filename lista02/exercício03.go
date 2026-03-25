package main

import "fmt"

func main() {
	var a, b, soma int
	fmt.Scan(&a, &b)
	soma = a + b
	if soma > 20 {
		soma = soma + 8
		fmt.Println(soma)
	} else {
		soma = soma - 5
		fmt.Println(soma)
	}
}
