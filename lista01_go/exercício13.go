package main
import "fmt"

func main() {
	var n float64
	fmt.Print("Digite a nota do aluno: ")
	fmt.Scan(&n)

	if n >= 9 && n <= 10 {
		fmt.Printf("Nota: %.1f Conceito: A\n", n)
	} else if n >= 7.5 && n < 9 {
		fmt.Printf("Nota: %.1f Conceito: B\n", n)
	} else if n >= 6 && n < 7.5 {
		fmt.Printf("Nota: %.1f Conceito: C\n", n)
	} else {
		fmt.Printf("Nota: %.1f Conceito: D\n", n)
	}
}
