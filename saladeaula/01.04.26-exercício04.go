package main

import f "fmt"

func fatorial(a int) int {
	var r int
	r = 1
	for i := a; i > 0; i-- {
		r = i * r
	}
	return r
}

func main() {
	var n int
	f.Scan(&n)
	f.Println(fatorial(n))
}
