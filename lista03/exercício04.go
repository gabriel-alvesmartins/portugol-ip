package main

import f "fmt"

func main() {
	var n int
	var r int
	for {
		f.Scan(&n)
		if n <= 0 {
			break
		} else {
			for i := 1; i < n; i++ {
				if i*i == n {
					r = 1
					break
				}
			}
			if r == 1 {
				f.Println("Sim")
				r = 0
			} else {
				f.Println("Não")
			}
		}
	}
}
