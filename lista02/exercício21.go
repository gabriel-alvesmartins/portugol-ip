package main

import (
	"fmt"
)

func main() {
	var (
		ni                         int
		n1, n2, n3, mediaE, mediaF float64
		conceito                   string
	)
	fmt.Println("Número de identificação do aluno")
	fmt.Scan(&ni)
	fmt.Println("Quais as notas desse aluno?")
	fmt.Scan(&n1, &n2, &n3)
	mediaE = (n1 + n2 + n3) / 3
	mediaF = (n1 + n2*2 + n3*3 + mediaE) / 7
	if mediaF >= 9 && mediaF <= 10 {
		conceito = "A"
	} else if mediaF >= 7.5 && mediaF < 9 {
		conceito = "B"
	} else if mediaF >= 6 && mediaF < 7.5 {
		conceito = "C"
	} else if mediaF >= 4 && mediaF < 6 {
		conceito = "D"
	} else {
		conceito = "E"
	}
	fmt.Printf("Número do aluno: %d  \nNota 1: %f\nNota 2: %f\nNota 3: %f\nMédia dos exercícios: %f\nMédia de aproveitamento: %f\nConceito: %s\n", ni, n1, n2, n3, mediaE, mediaF, conceito)
	switch conceito {
	case "A", "B", "C":
		fmt.Println("APROVADO")
	default:
		fmt.Println("REPROVADO")
	}
}
