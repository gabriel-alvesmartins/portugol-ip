package main

import f "fmt"

func main() {
	var n int
	f.Scan(&n)
	var fatorial int = 1
	if n<0 {
		f.Print("Número Inválido")
		return
	}
	for i:=n; i>=1; i--{
		fatorial *= i
	}
	f.Println(fatorial)
}
