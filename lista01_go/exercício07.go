package main
import "fmt"

func main() {
	var tempF, pol, tempC, mm float64

	fmt.Println("Dê um valor em temperatura Fahrenheit e outra em polegada:")
	fmt.Scan(&tempF, &pol)

	tempC = (5*tempF - 160) / 9
	mm = pol * 25.4

	fmt.Printf("O VALOR EM CELSIUS É: %.2f\n", tempC)
	fmt.Printf("A QUANTIDADE DE CHUVA É: %.2f\n", mm)
}
