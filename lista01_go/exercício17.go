package main
import "fmt"

func main() {
	var n, q int
	fmt.Println("Digite um número PAR e em sequência a QUANTIDADE de números que pares que o sucedam:")
	fmt.Scan(&n, &q)

	if n%2 == 0 {
		for i := 0; i <= (q*2)-2; i++ {
			if i%2 == 0 {
				valor := n + i
				fmt.Println(valor)
			}
		}
	} else {
		fmt.Println("O PRIMEIRO NÚMERO NÃO É PAR")
	}
}
