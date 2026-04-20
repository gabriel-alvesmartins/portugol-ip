package main
import "fmt"

func main() {
	var n1, n2 int
	fmt.Print("Digite N1 e N2: ")
	fmt.Scan(&n1, &n2)

	resultado := 0
	// Adiciona N1 a si mesmo N2 vezes
	for i := 0; i < n2; i++ {
		resultado += n1
	}
	fmt.Printf("A multiplicação de %d por %d é: %d\n", n1, n2, resultado)
}
