package main

import f "fmt"

func main() {
	var soma uint = 1
	var numero uint= 1
	for i:=1; i<=64; i++{
		numero *= 2
		soma += numero
	}
	f.Print(soma)
}
