package main

import f "fmt"

func main() {
	var n int;
	f.Scan(&n)
	var triangular bool = false
	for i:=1; i<n-3; i++{
		 if i*(i+1)*(i+2)==n{
			triangular = true
		 } 
	}
	if triangular==true{
		f.Println("É triangular")
	} else {
		f.Println("Não é triangular")
	}
}
