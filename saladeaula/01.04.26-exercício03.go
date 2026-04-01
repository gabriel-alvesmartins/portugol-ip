package main

import f "fmt"

func media(a, b, c float64) float64 {
	return (a + b + c) / float64(3)
}

func main() {
	var n1, n2, n3 float64
	f.Scan(&n1, &n2, &n3)
	m := media(n1, n2, n3)
	f.Println(m)
}
