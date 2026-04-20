package main
import "fmt"

func main() {
	soma := 0.0
	for i := 1.0; i <= 37.0; i++ {
		soma += ((38.0 - i + 1.0) * (38.0 - i)) / i
	}
	fmt.Printf("A soma S é: %.2f\n", soma)
}
