package main
import "fmt"

func main() {
	var horas int
	var valor float64

	fmt.Print("Quantas horas o cliente usou a charrete? ")
	fmt.Scan(&horas)

	valor = float64(((horas-horas%3)/3)*10 + (horas%3)*5)
	fmt.Printf("O VALOR A PAGAR É: %.2f\n", valor)
}
