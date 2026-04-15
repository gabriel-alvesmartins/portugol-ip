package main

import f "fmt"

func main() {
	var (
		n            = 0
		soma         = 0
		nd           = 0
		np           = 0
		ni           = 0
		sp           = 0
		maior = 0
		menor = 100000000000000000
	)

	var mnd, mdnp, pni float64
	for {
		f.Scan(&n)
		if n != 30000 {
			soma += n
			nd++
			if maior < n {
				maior = n
			}
			if menor > n {
				menor = n
			}
			if n%2 == 0 {
				sp += n
				np++
			} else {
				ni++
			}
		} else {
			break
		}
	}
	mnd = float64(soma) / float64(nd)
	mdnp = float64(sp) / float64(np)
	pni = (float64(ni) / float64(nd))*100

	f.Println(soma, nd, mnd, maior, menor, mdnp,pni)
}
