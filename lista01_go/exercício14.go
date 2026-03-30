package main
import (
	"fmt"
	"math"
)

func main() {
	var h, a, volume float64

	fmt.Println("Qual o valor da altura e da aresta da pirâmide, respectivamente? (em metros)")
	fmt.Scan(&h, &a)

	volume = 3 * math.Pow(a, 2) * math.Sqrt(3) * h / 6
	fmt.Printf("O VOLUME DA PIRAMIDE É: %.2f METROS CUBICOS\n", volume)
}
