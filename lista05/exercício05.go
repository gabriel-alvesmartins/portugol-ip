package main

import f "fmt"

func main() {
	var v[10]int;

	
	for i:=0; i<10; i++{
		f.Scan(&v[i])
	}
	menor:= v[0]

	for i:=0; i<10; i++{
		if menor>v[i]{
			menor = v[i]
		}
	}
	for i:=0; i<10; i++{
		if menor == v[i]{
			f.Printf("O menor valor e sua posição no vetor são respectivamente: %d e %d",menor, i )
		}
	}
}
