package main
import "fmt"

func fatorial(n float64) float64 {
	if n == 0 { return 1 }
	fat := 1.0
	for i := 1.0; i <= n; i++ { fat *= i }
	return fat
}

func main() {
	soma := 0.0
	numerador := 100.0
	
	for i := 0.0; i < 20.0; i++ {
		soma += numerador / fatorial(i)
		numerador--
	}
	fmt.Printf("A soma é: %f\n", soma)
}
