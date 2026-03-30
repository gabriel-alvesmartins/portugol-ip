package main

import "fmt"

func main() {
	var (
		n    [5]float64
		soma float64
	)
	for i := 0; i < len(n); i++ {
		fmt.Scan(&n[i])
		soma += n[i]
	}
	fmt.Println(soma)
}
