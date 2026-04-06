package main

import f "fmt"

type Pessoa struct {
	nome   string
	altura float64
	peso   float64
}

func main() {
	var v []Pessoa
	var p Pessoa
	for {
		f.Scan(&p.nome)
		if p.nome == "FIM" {
			break
		}
		f.Scan(&p.altura)
		p.peso = (72.7 * p.altura) - 58.0
		v = append(v, p)
	}
	f.Print(v)
}
