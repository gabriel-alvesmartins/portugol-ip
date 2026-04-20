package main
import "fmt"

func main() {
	fmt.Println("Angulo(A)\tSeno(A)")
	for a := 0.0; a <= 6.301; a += 0.1 { 
		a3 := a * a * a
		a5 := a3 * a * a
		a7 := a5 * a * a
		senA := a - (a3 / 6.0) + (a5 / 120.0) - (a7 / 5040.0)
		fmt.Printf("%.1f\t\t%f\n", a, senA)
	}
}
