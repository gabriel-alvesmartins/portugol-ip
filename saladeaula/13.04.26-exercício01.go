package main

import f "fmt"

func busca(v []int, x int) int {
	for i := 0; i < len(v); i++ {
		if v[i] == x {
			return i
		}
	}
	return -1
}

func main() {
	var (
		n int
		v []int
	)
	f.Scan(&n)
	v = make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		f.Scan(&a)
		v[i] = a
	}
	var x int
	f.Scan(&x)
	resultado := busca(v, x)
	f.Print(resultado, "\n")
}
