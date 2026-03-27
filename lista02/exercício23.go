package main

import (
	"fmt"
)

func main() {
	var i uint
	fmt.Scan(&i)
	if i < 16 {
		fmt.Println("Não-eleitor")
	} else if i >= 18 && i <= 65 {
		fmt.Println("Eleitor Obrigatório")
	} else {
		fmt.Println("Eleitor Facultativo")
	}
}
