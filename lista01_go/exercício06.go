package main
import "fmt"

func main() {
	var n int
	fmt.Print("Quantas temperaturas você quer converter de Fahrenheit para Celsius? ")
	fmt.Scan(&n)

	tempF := make([]float64, n)
	tempC := make([]float64, n)

	for i := 0; i < n; i++ {
		fmt.Printf("Qual o valor da %dª temperatura? ", i+1)
		fmt.Scan(&tempF[i])
		tempC[i] = 5 * (tempF[i] - 32) / 9
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%.2fº FAHRENHEIT EQUIVALE A %.2fº CELSIUS\n", tempF[i], tempC[i])
	}
}
