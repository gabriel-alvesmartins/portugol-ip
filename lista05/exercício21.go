package main

import f "fmt"

func main() {
	var v [10]float64
	var codigo int
	for i := 0; i < 10; i++ {
		f.Scan(&v[i])
	}

	for {
		f.Println("\nDigite o código (0 para Sair, 1 para Direta, 2 para Inversa):")
		f.Scan(&codigo)

		if codigo == 0 {
			
			break 
		} else if codigo == 1 {
			
			for i := 0; i < 10; i++ {
				f.Print(v[i], " ")
			}
			f.Println()
		} else if codigo == 2 {
			
			for i := 9; i >= 0; i-- {
				f.Print(v[i], " ")
			}
			f.Println()
		} else {
			f.Println("Código inválido!")
		}
	}
}
