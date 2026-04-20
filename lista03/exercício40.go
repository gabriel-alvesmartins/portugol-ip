package main

import (
	"fmt"
)

func main() {
	despesas := 300.0
	ingressos := 130

	var lucroMaximo float64
	var precoMelhor float64
	var ingressosMelhor int
	primeiro := true 

	fmt.Println("Preço(R$)\tIngressos\tLucro(R$)")
	fmt.Println("-----------------------------------------")

	for preco := 6.0; preco > 0.99; preco -= 0.60 {
		lucro := (preco * float64(ingressos)) - despesas

		fmt.Printf("%.2f\t\t%d\t\t%.2f\n", preco, ingressos, lucro)

		
		if primeiro || lucro > lucroMaximo {
			lucroMaximo = lucro
			precoMelhor = preco
			ingressosMelhor = ingressos
			primeiro = false
		}

		
		ingressos += 30
	}

	fmt.Println("\n--- Resultado Ideal ---")
	fmt.Printf("Lucro Máximo Esperado: R$ %.2f\n", lucroMaximo)
	fmt.Printf("Preço ideal do Ingresso: R$ %.2f\n", precoMelhor)
	fmt.Printf("Quantidade de Ingressos esperada: %d\n", ingressosMelhor)
}
