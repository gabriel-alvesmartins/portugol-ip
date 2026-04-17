package main

import f "fmt"

func main() {
	var n int
	f.Scan(&n)

	v := make([][2]float64, n)

	var (
		Aprovado  = 0
		Exame     = 0
		Reprovado = 0
	)
	var somat float64 = 0
	var soma2 float64 = 0
	var medialuno float64 = 0
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			f.Scan(&v[i][j])
			if j == 0 {
				soma2 += v[i][j]
				continue
			} else {
				soma2 += v[i][j]
				medialuno = soma2 / float64(2)
				somat += medialuno
				f.Println("Media aritimética do ", i+1, "º aluno:", medialuno)
				if medialuno > 7 {
					f.Println("Aprovado")
					Aprovado++
				} else if medialuno >= 3 && medialuno <= 7 {
					f.Println("Exame")
					Exame++
				} else {
					f.Println("Reprovado")
					Reprovado++
				}
				soma2 = 0
			}
		}
	}
	var media float64 = somat / float64(n)
	f.Println(Aprovado, Exame, Reprovado, media)
}
