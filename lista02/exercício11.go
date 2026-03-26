package main

import "fmt"

func main() {
	var n int
	var v float64
	fmt.Scan(&n)
	v = 8 / float64(2-n)
	fmt.Println(v)
}
