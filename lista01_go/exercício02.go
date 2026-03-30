package main
import "fmt"

func main() {
	var nj int
	var total, pop, ge, arq, cad float64

	fmt.Print("Quantos jogos você quer analisar? ")
	fmt.Scan(&nj)

	arrecadacao := make([]float64, nj) // Usamos slice no lugar do vetor gigante

	for i := 0; i < nj; i++ {
		fmt.Println("Responda os seguintes questionamentos:")
		fmt.Println("Qual a quantidade TOTAL, e as % de POPULAR, GERAL, ARQUIBANCADA e CADEIRAS?")
		fmt.Scan(&total, &pop, &ge, &arq, &cad)
		
		arrecadacao[i] = ((total * pop / 100) * 1) + ((total * ge / 100) * 5) + ((total * arq / 100) * 10) + ((total * cad / 100) * 20)
	}

	for i := 0; i < nj; i++ {
		fmt.Printf("A RENDA DO JOGO Nº %d É %.2f\n", i+1, arrecadacao[i])
	}
}
