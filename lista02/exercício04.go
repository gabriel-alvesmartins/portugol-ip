package main

import (
	"fmt"
	"math"
)

func main() {
	var n float64
	fmt.Scan(&n)
	var resultado float64
	if n < 0 {
		resultado = n * n
		fmt.Println(resultado)
	} else {
		resultado = math.Sqrt(n)
		fmt.Println(resultado)
	}
}
