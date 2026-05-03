package main

import f "fmt"

func main() {
	var codigos [10]int
	var saldos [10]float64

	f.Println("--- Cadastro do Banco ---")
	for i := 0; i < 10; i++ {
		f.Printf("Código da Conta %d: ", i+1)
		f.Scan(&codigos[i])
		f.Printf("Saldo da Conta %d: ", i+1)
		f.Scan(&saldos[i])
	}

	var opcao int
	for {
		f.Println("\n--- Menu ---")
		f.Println("1. Efetuar depósito")
		f.Println("2. Efetuar saque")
		f.Println("3. Consultar ativo bancário")
		f.Println("4. Finalizar programa")
		f.Print("Opção: ")
		f.Scan(&opcao)

		if opcao == 4 {
			break 
		} else if opcao == 1 {
			var codBusca int
			var valor float64
			f.Print("Código da conta: ")
			f.Scan(&codBusca)

			encontrou := false
			for i := 0; i < 10; i++ {
				if codigos[i] == codBusca {
					f.Print("Valor do depósito: ")
					f.Scan(&valor)
					saldos[i] += valor
					f.Println("Depósito efetuado! Novo saldo:", saldos[i])
					encontrou = true
					break
				}
			}
			if !encontrou {
				f.Println("Conta não encontrada.")
			}

		} else if opcao == 2 {
			var codBusca int
			var valor float64
			f.Print("Código da conta: ")
			f.Scan(&codBusca)

			encontrou := false
			for i := 0; i < 10; i++ {
				if codigos[i] == codBusca {
					f.Print("Valor do saque: ")
					f.Scan(&valor)
					if saldos[i] >= valor {
						saldos[i] -= valor
						f.Println("Saque efetuado! Novo saldo:", saldos[i])
					} else {
						f.Println("Saldo insuficiente.")
					}
					encontrou = true
					break
				}
			}
			if !encontrou {
				f.Println("Conta não encontrada.")
			}

		} else if opcao == 3 {
			var soma float64 = 0
			for i := 0; i < 10; i++ {
				soma += saldos[i]
			}
			f.Printf("O ativo bancário é de: R$ %.2f\n", soma)
		} else {
			f.Println("Opção inválida.")
		}
	}
}
