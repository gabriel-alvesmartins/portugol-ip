package main
import "fmt"

func main() {
	var a, b, c, delta float64

	fmt.Println("Digite os 3 coeficientes, em ordem (a,b,c), para calcular o delta:")
	fmt.Scan(&a, &b, &c)

	delta = (b * b) - (4 * a * c)

	fmt.Printf("O VALOR DE DELTA É %.2f\n", delta)
}
