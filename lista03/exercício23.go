package main
import "fmt"

func main() {
	var n int
	fmt.Print("Digite N: ")
	fmt.Scan(&n)

	soma := 0.0
	sinal := 1.0
	numerador := 1000.0

	for i := 1.0; i <= float64(n); i++ {
		soma += sinal * (numerador / i)
		numerador -= 3.0
		sinal *= -1.0 
	}
	fmt.Printf("O resultado é: %.2f\n", soma)
}
