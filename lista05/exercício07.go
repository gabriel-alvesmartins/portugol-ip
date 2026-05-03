package main

import f "fmt"

func main() {
	var v[]int;

	for i:=1; i<200; i+=2{
		v = append(v, i)
	}
	f.Print(v)
	
}
