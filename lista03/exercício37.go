package main

import f "fmt"

func main() {
	var n string
	f.Scan(&n)

	var (
		numero int
		multiplicador = 1
		soma = 0
	)

	for i:=len(n)-1; i>=0; i--{
		numero = multiplicador * int(n[i] - '0')
		soma += numero
		multiplicador *=8
	}
	f.Print(soma)
}
