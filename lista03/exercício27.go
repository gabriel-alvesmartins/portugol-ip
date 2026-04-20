package main

import (
	"fmt"
	"math"
)

func fatorial(n float64) float64 {
	if n == 0 { return 1 }
	fat := 1.0
	for i := 1.0; i <= n; i++ { fat *= i }
	return fat
}

func main() {
	var x float64
	fmt.Print("Digite x: ")
	fmt.Scan(&x)

	cosCalculado := 0.0
	sinal := 1.0

	for i := 0.0; i < 20.0; i++ {
		expoente := 2.0 * i
		termo := math.Pow(x, expoente) / fatorial(expoente)
		cosCalculado += sinal * termo
		sinal *= -1.0
	}

	cosReal := math.Cos(x)
	diferenca := math.Abs(cosCalculado - cosReal)

	fmt.Printf("Cosseno Calculado (série): %f\n", cosCalculado)
	fmt.Printf("Cosseno Função (math.Cos): %f\n", cosReal)
	fmt.Printf("Diferença: %f\n", diferenca)
}
