package main

import (
	"fmt"
)

func main() {
	var (
		mf, he          int
		she, sb, ir, sl float64
	)
	const (
		sm  = 788
		vhe = 10
	)
	fmt.Println("Digite a matrícula do funcionário e a quantidade de horas trabalhadas:")
	fmt.Scan(&mf, &he)
	she = float64(he) * vhe
	sb = 3*sm + she
	ir = sb * 0.2
	sl = sb - ir
	fmt.Printf("Salário Bruto: %f\nSalário Líquido: %f\nImposto de Renda: %f\nSalário hora-extra: %f\n", sb, sl, ir, she)
}
