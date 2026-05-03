package main

import f "fmt"

func main() {

	var empregados [100]int
	var meses [100]int
	var quantidade int = 0 

	f.Println("Digite o número do empregado e os meses trabalhados.")
	f.Println("Para encerrar, digite 0 0:")

	
	for i := 0; i < 100; i++ {
		var id, mes int
		f.Scan(&id, &mes)

		
		if id == 0 && mes == 0 {
			break
		}

		empregados[i] = id
		meses[i] = mes
		quantidade++
	}

	if quantidade == 0 {
		f.Println("Nenhum empregado foi cadastrado.")
		return
	}

	qtdProcurar := 3
	if quantidade < 3 {
		qtdProcurar = quantidade
	}

	f.Println("\n--- Os empregados mais recentes (menos meses) ---")

	
	for k := 0; k < qtdProcurar; k++ {
		menorMes := 999999 
		indiceMenor := -1

		
		for i := 0; i < quantidade; i++ {
			if meses[i] < menorMes {
				menorMes = meses[i]
				indiceMenor = i
			}
		}

		
		f.Printf("%dº mais recente: Empregado nº %d (Trabalha há %d meses)\n", k+1, empregados[indiceMenor], menorMes)

		meses[indiceMenor] = 999999
	}
}
