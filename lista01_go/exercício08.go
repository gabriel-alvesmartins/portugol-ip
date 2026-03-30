package main
import (
	"fmt"
	"math"
)

func main() {
	var areaT, custo, r, h float64

	fmt.Println("Qual o raio e a altura da lata, respectivamente? (em metros)")
	fmt.Scan(&r, &h)

	areaT = 2*math.Pi*math.Pow(r, 2) + 2*math.Pi*r*h
	custo = areaT * 100

	fmt.Printf("O VALOR DO CUSTO É: %.2f\n", custo)
}
