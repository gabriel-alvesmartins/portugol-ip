package main
import "fmt"

func main() {
	var n int
	var soma float64

	fmt.Print("Digite um algarismo inteiro maior que 1: ")
	fmt.Scan(&n)

	soma = 0
	if n > 1 {
		for i := 1; i <= n; i++ {
			valor := 1.0 / float64(i)
			soma = soma + valor
		}
		fmt.Printf("%.6f\n", soma)
	} else {
		fmt.Println("Numero invalido!")
	}
}
