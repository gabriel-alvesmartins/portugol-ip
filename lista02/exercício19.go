package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		f          int
		r, h, a, v float64
	)
	const Pi = 3.14159265358979323846264338327950288419716939937510582097494459
	fmt.Println("Selecione uma figura 1- cone 2- cilindro 3- esfera")
	fmt.Scan(&f)
	switch f {
	case 1:
		fmt.Println("Dê os valores de raio e altura: ")
		fmt.Scan(&r, &h)
		v = (Pi * r * r * h) / 3
		a = (Pi * r * math.Sqrt((r*r + h*h)))
		fmt.Println("Volume: ", v, "\nÁrea: ", a)
	case 2:
		fmt.Println("Dê os valores de raio e altura: ")
		fmt.Scan(&r, &h)
		v = (Pi * r * r * h)
		a = (2 * Pi * r * h)
		fmt.Println("Volume: ", v, "\nÁrea: ", a)
	case 3:
		h = 0
		fmt.Println("Dê o valor do raio: ")
		v = (4 / 3) * Pi * math.Pow(r, 3)
		a = 4 * Pi * r * r
		fmt.Println("Volume: ", v, "\nÁrea: ", a)
	default:
		fmt.Println("Escolha uma figura válida!")
	}
}
