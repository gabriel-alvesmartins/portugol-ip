package main

import f "fmt"

func main() {
	var v[]int;

	for i:=100; i>0; i--{
		v = append(v, i)
	}
	f.Print(v)
	
}
