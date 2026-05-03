package main

import f "fmt"

func main() {
	var janela [24]int
	var corredor [24]int

	var tipo int
	for {
		cheioJanela := true
		cheioCorredor := true
		for i := 0; i < 24; i++ {
			if janela[i] == 0 {
				cheioJanela = false
			}
			if corredor[i] == 0 {
				cheioCorredor = false
			}
		}

		if cheioJanela && cheioCorredor {
			f.Println("Ônibus completamente cheio!")
			break 
		}

		f.Println("\n--- Venda de Passagens ---")
		f.Println("1 para Janela, 2 para Corredor, ou 0 para Sair")
		f.Print("Sua escolha: ")
		f.Scan(&tipo)

		if tipo == 0 {
			break
		} else if tipo == 1 {
			if cheioJanela {
				f.Println("Não há poltronas livres na janela.")
			} else {
				f.Print("Poltronas livres na Janela: ")
				for i := 0; i < 24; i++ {
					if janela[i] == 0 {
						f.Print(i+1, " ")
					}
				}
				f.Println("\nQual poltrona deseja comprar?")
				var poltrona int
				f.Scan(&poltrona)

				
				if poltrona >= 1 && poltrona <= 24 && janela[poltrona-1] == 0 {
					janela[poltrona-1] = 1 
					f.Println("Venda na janela realizada com sucesso!")
				} else {
					f.Println("Poltrona inválida ou já ocupada.")
				}
			}

		} else if tipo == 2 {
			if cheioCorredor {
				f.Println("Não há poltronas livres no corredor.")
			} else {
				f.Print("Poltronas livres no Corredor: ")
				for i := 0; i < 24; i++ {
					if corredor[i] == 0 {
						f.Print(i+1, " ")
					}
				}
				f.Println("\nQual poltrona deseja comprar?")
				var poltrona int
				f.Scan(&poltrona)

				if poltrona >= 1 && poltrona <= 24 && corredor[poltrona-1] == 0 {
					corredor[poltrona-1] = 1 
					f.Println("Venda no corredor realizada com sucesso!")
				} else {
					f.Println("Poltrona inválida ou já ocupada.")
				}
			}
		} else {
			f.Println("Opção inválida!")
		}
	}
}
