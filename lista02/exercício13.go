package main

import "fmt"

func main() {
	var v, n uint
	fmt.Scan(&n)
	if n >= 100 && n <= 999 {
		v = ((n % 100) - (n%100)%10) / 10
		fmt.Println(v)
	} else {
		fmt.Println("Número inválido")
	}
}
