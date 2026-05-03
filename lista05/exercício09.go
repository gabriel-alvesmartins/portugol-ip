package main

import f "fmt"

func main() {
	var v[10]float64;

	var soma float64 = 0

	for i:=0; i<10 ; i++{
		f.Scan(&v[i])
		soma+=v[i]
	}

	media:=soma/10

	for i:=0; i<10; i++{
		if v[i]> media{
			f.Print(v[i], ",")
		}
	}
	
}
