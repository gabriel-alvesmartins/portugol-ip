package main

import (
	"fmt"
)
func main() {
	var n int
	fmt.Scan(&n)
	if n > 20 && n < 90 {
		fmt.Println("Seu número está entre o intervalo 20<n<90")
	} else {
		fmt.Println("Seu número não está entre o intervalo 20<n<90")
	}
}
