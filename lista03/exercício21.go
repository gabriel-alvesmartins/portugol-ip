package main

import f "fmt"

func main() {
	var a, p, r int
	f.Scan(&a, &p)
	r = a
	for i := 0; i < p-1; i++ {
		a = a * r
	}
	f.Println(a)
}
