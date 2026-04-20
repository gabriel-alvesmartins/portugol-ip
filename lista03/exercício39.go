package main

import (
	"fmt"
)

func main() {
	var idMaisGordo, idMaisMagro int
	var pesoMaisGordo, pesoMaisMagro float64

	for i := 1; i <= 90; i++ {
		var id int
		var peso float64

		fmt.Printf("Digite o ID e o peso do boi %d (separados por espaço): ", i)
		fmt.Scan(&id, &peso)

		if i == 1 {
			idMaisGordo = id
			pesoMaisGordo = peso
			idMaisMagro = id
			pesoMaisMagro = peso
		} else {
			
			if peso > pesoMaisGordo {
				pesoMaisGordo = peso
				idMaisGordo = id
			}

			if peso < pesoMaisMagro {
				pesoMaisMagro = peso
				idMaisMagro = id
			}
		}
	}

	fmt.Println("\n--- Resultados ---")
	fmt.Printf("Boi mais gordo: ID %d com peso de %.2f kg\n", idMaisGordo, pesoMaisGordo)
	fmt.Printf("Boi mais magro: ID %d com peso de %.2f kg\n", idMaisMagro, pesoMaisMagro)
}
