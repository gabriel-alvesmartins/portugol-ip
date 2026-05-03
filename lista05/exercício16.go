package main

import f "fmt"

func main() {

	var v [50]int
	var jacontei [50]bool

	for i := 0; i < 50; i++ {
		f.Scan(&v[i])
	}

	moda := 0
	var numeromoda int

	for i := 0; i < 50; i++ {
		if jacontei[i] {
			continue
		}
		contador := 0
		for j := i + 1; j < 50; j++ {
			if v[i] == v[j] {
				contador++
			}
		}
		if contador > moda {
			moda = contador
			numeromoda = v[i]
		}
	}

	f.Print(numeromoda, ": ", moda, " repetições")
}
