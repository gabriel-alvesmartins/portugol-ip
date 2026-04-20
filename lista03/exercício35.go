package main

import f "fmt"

func main() {
	var n int
	f.Scan(&n)
	var v []int

	for n != 0 {
		if n%2 == 0 {
			v=append(v,0)
		} else {
			v=append(v,1)
		}
		n = n / 2
	}
	for i:=len(v)-1; i>=0; i--{
		f.Print(v[i])
	}

}
