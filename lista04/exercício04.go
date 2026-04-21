package main

import "fmt"

func decimalParaBinario(n int) {
	
	if n/2 > 0 {
		decimalParaBinario(n / 2)
	}
	
	
	fmt.Print(n % 2)
}

func main() {
	numero := 10
	fmt.Printf("O número %d em binário é: ", numero)
	decimalParaBinario(numero)
	fmt.Println() 
}
