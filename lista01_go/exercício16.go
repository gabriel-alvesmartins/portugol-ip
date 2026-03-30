package main
import "fmt"

func main() {
	var sal, nvsal float64

	fmt.Println("Escreva o salário do funcionário:")
	fmt.Scan(&sal)

	if sal <= 300 {
		nvsal = 1.5 * sal
	} else {
		nvsal = 1.3 * sal
	}
	fmt.Printf("SALARIO COM REAJUSTE: %.2f\n", nvsal)
}
