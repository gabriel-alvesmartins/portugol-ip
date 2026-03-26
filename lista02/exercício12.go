package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Insira uma idade válida")
	} else {
		if n <= 2 {
			fmt.Println("Recém-nascido")
		} else if n >= 3 && n <= 11 {
			fmt.Println("Criança")
		} else if n >= 12 && n <= 19 {
			fmt.Println("Adolescente")
		} else if n >= 20 && n <= 55 {
			fmt.Println("Adulto")
		} else {
			fmt.Println("Adulto")
		}
	}
}
