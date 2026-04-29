package main

import f "fmt"

func main() {
	var v [10]int;
	verificador:= false;
	for i:=0 ; i<10; i++{
		var a int;
		f.Scan(&a);
		v[i] = a;

		if (v[i]>50){
			verificador = true
			f.Printf("Número: %d Posição: %d \n", v[i], i+1)
		} 
	}
	if verificador==false{
		f.Print("-1 \n");
	}
}
