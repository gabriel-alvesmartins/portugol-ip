package main

import f "fmt"

var v []pessoa

type pessoa struct {
	i int
	p float64
	a float64
}

func main() {
	var p pessoa
	for {
		f.Println("Digite 1 se você quer continuar e 2 se quer parar de adicionar dados")
		var r int
		f.Scan(&r)
		if r == 1 {
			f.Scan(&p.i, &p.a, &p.p)
			v = append(v, p)
		} else {
			break
		}
	}
	var p50 int = 0
	var sh float64 = 0
	var pesot float64 = 0
	var m40kg int = 0
	var t1020 int = 0

	for i := 0; i < len(v); i++ {

		pesot = pesot + v[i].p

		if v[i].p < 40 {
			m40kg++
		}

		if v[i].i >= 10 && v[i].i <= 20 {
			sh = sh + v[i].a
			t1020++
		}

		if v[i].i > 50 {
			p50++
		}
	}
	var media  float64= 0;
	if t1020 !=0{
		media = sh / float64(t1020)
	}
	var porcentagem float64 = float64(m40kg) /float64(len(v))*100
	
	f.Println(p50,"\n",media,"\n",porcentagem)
}
