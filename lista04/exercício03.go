package main

import "fmt"


func inverteArray(arr []int, n int) {
	
	if n <= 1 {
		return
	}

	arr[0], arr[n-1] = arr[n-1], arr[0]

	inverteArray(arr[1:n-1], n-2)
}

func main() {
	numeros := []int{10, 20, 30, 40, 50}
	tamanho := len(numeros)

	fmt.Println("Array Original:", numeros)

	inverteArray(numeros, tamanho)

	fmt.Println("Array Invertido:", numeros)
}
