package main

import f "fmt"

func main() {
	var soma float64 = 0;
	var n float64 = 1;
	for i:=1;i<=50;i++{
		soma+=n/float64(i)
		n+=2
	}
	f.Print(soma)
}
