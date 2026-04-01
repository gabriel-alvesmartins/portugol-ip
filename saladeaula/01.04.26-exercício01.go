package main

import f "fmt"

func soma(a, b float64) float64 {
	r := a + b
	return r
}

func main() {
	var n1, n2 float64
	f.Scan(&n1, &n2)
	resultado := soma(n1, n2)
	f.Println(resultado)
}

