package main

import "fmt"

func main() {
	var d, m, a int
	var mes string
	fmt.Scan(&d, &m, &a)
	switch m {
	case 1:
		mes = "Janeiro"
	case 2:
		mes = "Fevereiro"
	case 3:
		mes = "Março"
	case 4:
		mes = "Abril"
	case 5:
		mes = "Maio"
	case 6:
		mes = "Junho"
	case 7:
		mes = "Julho"
	case 8:
		mes = "Agosto"
	case 9:
		mes = "Setembro"
	case 10:
		mes = "Outubro"
	case 11:
		mes = "Novembro"
	case 12:
		mes = "Dezembro"
	default:
		fmt.Println("Insira um mês válido entre 1-12")
	}
	fmt.Printf("A sua data é: %d de %s de %d", d, mes, a)
}
