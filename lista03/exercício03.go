package main

import f "fmt"

func main() {
	var sc, sj float64
	var mes int
	f.Scan(&sc)
	sj = sc / 3
	mes = 0
	for {
		sc = sc * 1.02
		sj = sj * 1.05
		mes++
		if sj >= sc {
			break
		}
	}
	f.Print(mes, "\n")
}
