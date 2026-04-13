package main

import f "fmt"

func main() {
	var soma = 0
	for i := 50; i <= 70; i++ {
		soma += i
	}
	var media float64
	media = float64(soma) / 20
	f.Print("Soma: ", soma, "\n")
	f.Print("Media: ", media, "\n")
}
