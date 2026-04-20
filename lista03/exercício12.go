package main

import f "fmt"

func main() {
	var x float64
	f.Scan(&x)
	var fatorial float64 = 1
	var soma float64 = 0;
	for i:=1; i<=19; i++{
		fatorial = fatorial * float64(i)
		if i%2==0{
			soma+=(x/fatorial)
		} else {
			soma-=(x/fatorial)
		}
		
	}
	var somatorio float64 = (x + soma)
	f.Print(somatorio)
}
