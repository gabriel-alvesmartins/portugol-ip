package main

import f "fmt"

func main() {
	
	v := make([]int, 11)

	for i:=0; i<11; i++{
		f.Scan(&v[i])
	}

	var soma1 int = 0;
	var n int = 10;
	for i:=0; i<9; i++{
		soma1 += v[i]*n
		n--;
	}

	var dv1 int
	if soma1%11<2{
		dv1 = 0
	} else{
		dv1 = 11-(soma1%11)
	}

	var soma2 int = 0;
	var m int = 11;
	for i:=0; i<10; i++{
		soma2 += v[i]*m
		m--;
	}

	var dv2 int
	if soma2%11<2{
		dv2 = 0
	} else{
		dv2 = 11-(soma2%11)
	}

	if dv1==v[9] && dv2==v[10]{
		f.Println("Válido")
	} else {
		f.Println("Inválido")
	}
}
