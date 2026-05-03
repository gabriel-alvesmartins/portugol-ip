package main

import f "fmt"

func main() {
	var v1[10]int;
	var v2[5]int;
	var vr1[]int
	var vr2[]int

	for i:= 0; i<10; i++{
		var a int
		f.Scan(&a);
		v1[i] = a;
	}

	var soma int= 0;

	for i:= 0; i<5; i++{
		var a int
		f.Scan(&a);
		v2[i] = a;
		soma+=a;
	}

	for i:= 0; i<10; i++{
		if v1[i]%2==0{
			b:= soma+v1[i]
			vr1 = append(vr1, b)
		} else {
			b:= soma+v1[i]
			vr2 = append(vr2, b)
		}
	}

	f.Printf("Primeiro Vetor: %d\nSegundo Vetor: %d\nPrimeiro vetor resultante: %d\nSegundo vetor resultante: %d\n", v1,v2,vr1,vr2)
}
