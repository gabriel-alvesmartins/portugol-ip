package main

import f "fmt"

func main() {
	var v[50]float64;

	v[0] = 1
	v[1] = 1

	for i:=2; i<50; i++{
		v[i] = v[i-1] + v[i-2]
	}
	f.Print(v)
}
