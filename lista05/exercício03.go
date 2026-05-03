package main

import f "fmt"

func main() {
	var v[10]int;
	var pares[]int;
	var impares[]int;

	cnti:=0;
	soma:= 0;
	
	for i:=0 ; i<10; i++{
		var a int;
		f.Scan(&a)
		v[i] = a

		if a%2==0{
			pares = append(pares, a)
			soma+=a;
		} else{
			impares = append(impares, a)
			cnti++;
		}
	}

	f.Print(pares, "\n",soma, "\n",impares, "\n",cnti)
}
