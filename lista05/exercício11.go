package main

import f "fmt"

func main() {
	var v[100]float64;

	for i:=0; i<100; i++{
		f.Scan(&v[i])
	}

	var somatorio float64 = 0

	for i:=0; i<50; i++{
		somatorio+= (v[i]-v[99-i])*(v[i]-v[99-i])*(v[i]-v[99-i])
	}

	f.Print(somatorio)
}
