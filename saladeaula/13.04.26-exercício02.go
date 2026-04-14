package main

import (
	f "fmt"
	"sort"
)

func buscabinaria(v []int, x int) int {
	e := 0
	var d int = 0
	if len(v)%2 == 0 {
		d = len(v) - 1
	} else {
		d = len(v)
	}
	for e <= d {
		m := (e + d) / 2
		if v[m] == x {
			return m
		}
		if v[m] < x {
			e = m + 1
		}
		if v[m] > x {
			d = m - 1
		}
	}
	return -1
}

func main() {
	// primeiro digitar o tamanho do vetor, depois seus respectivos números. Após isso digitar o número a ser buscado no vetor.
	// O números do slice pode ser em qualquer ordem
	var v []int
	var n int
	f.Scan(&n)
	v = make([]int, n)
	for i := 0; i < n; i++ {
		var a int
		f.Scan(&a)
		v[i] = a
	}
	sort.Ints(v)
	x := 0
	f.Scan(&x)
	r := buscabinaria(v, x)
	f.Print(r, "\n")
}
