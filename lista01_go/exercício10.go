package main
import "fmt"

func main() {
	var matr [2][2]float64
	var determinante float64

	fmt.Println("Este é um calculador de matriz 2x2")
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Digite o %dº número da matriz da linha %d\n", j+1, i+1)
			fmt.Scan(&matr[i][j])
		}
	}

	determinante = matr[0][0]*matr[1][1] - matr[0][1]*matr[1][0]
	fmt.Println("O VALOR DO DETERMINANTE É:", determinante)
}
