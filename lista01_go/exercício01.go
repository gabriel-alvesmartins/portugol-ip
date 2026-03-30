package main
import "fmt"

func main() {
	var n1, n2, n3, media float64
	fmt.Println("Quais foram as notas do aluno?")
	fmt.Scan(&n1, &n2, &n3)
	
	media = (n1 + n2 + n3) / 3
	if media >= 6 {
		fmt.Println("Média =", media)
		fmt.Println("APROVADO")
	} else {
		fmt.Println("Média =", media)
		fmt.Println("REPROVADO")
	}
}
