package main
import "fmt"

func main() {
	soma := 0.0
	numerador := 1.0
	sinal := 1.0

	for i := 15.0; i >= 1.0; i-- {
		denominador := i * i
		soma += sinal * (numerador / denominador)
		numerador *= 2.0
		sinal *= -1.0
	}
	fmt.Printf("A soma S é: %.5f\n", soma)
}
