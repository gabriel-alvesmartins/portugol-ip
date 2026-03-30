package main

import "fmt"

func main() {
	var (
		n [10]float64
	)
	for i := 0; i < len(n); i++ {
		fmt.Scan(&n[i])
	}
	for j := 9; j == 0; j-- {
		fmt.Print(n[j], " ")
	}
}
