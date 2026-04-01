package main

import f "fmt"

func maior(a, b, c int) int {
	var maior int

	if b > a && b > c {
		maior = b
	} else if c > a && c > b {
		maior = c
	} else {
		maior = a
	}

	return maior
}

func main() {
	var n1, n2, n3 int
	f.Scan(&n1, &n2, &n3)
	m := maior(n1, n2, n3)
	f.Println(m)
}
