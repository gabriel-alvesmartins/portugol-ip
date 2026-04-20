package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Raio(R)\tVolume")
	for r := 0.0; r <= 20.0; r += 0.5 {
		volume := (4.0 / 3.0) * math.Pi * math.Pow(r, 3)
		fmt.Printf("%.1f\t%.2f\n", r, volume)
	}
}
