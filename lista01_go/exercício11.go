package main
import "fmt"

func main() {
	var n int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&n)

	if n%3 == 0 && n%5 == 0 {
		fmt.Println("O número é divisível por 3 e 5")
	} else {
		fmt.Println("O número não é divisível por 3 e 5")
	}
}
