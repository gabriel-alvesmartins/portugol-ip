package main
import "fmt"

func main() {
	soma := 0.0
	numerador := 1.0
	sinal := 1.0

	// O denominador começa em 15^2 (225) e vai até 1^2 (1), logo são 15 termos.
	for i := 15.0; i >= 1.0; i-- {
		denominador := i * i
		soma += sinal * (numerador / denominador)
		numerador *= 2.0
		sinal *= -1.0
	}
	fmt.Printf("A soma S é: %.5f\n", soma)
}
