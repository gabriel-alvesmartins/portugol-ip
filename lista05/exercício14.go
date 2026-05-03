package main

import f "fmt"

func main() {

	var a[10]int
	var b[10]int
	var c[]int

	for i:=0; i<10; i++{
		f.Scan(&a[i])
	}
	for i:=0; i<10; i++{
		f.Scan(&b[i])
	}

	for i:=0; i<10; i++{
		c = append(c, a[i], b[i])
	}
	f.Print(c)
}
