package main

import f "fmt"

func main() {
	var jogadas [20]int
	var freq [7]int 

	for i := 0; i < 20; i++ {
		f.Scan(&jogadas[i])

		if jogadas[i] >= 1 && jogadas[i] <= 6 {
			freq[jogadas[i]]++
		}
	}

	f.Println("Números sorteados:", jogadas)

	for i := 1; i <= 6; i++ {
		f.Printf("Face %d saiu %d vezes\n", i, freq[i])
	}
}
