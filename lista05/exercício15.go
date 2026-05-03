package main

import f "fmt"

func main() {

	var v[30]int
	var p[]int

	for i:=0;i<30;i++{
		f.Scan(&v[i])
		if i%2==0{
			p = append(p,v[i]*2)
		} else{
			p = append(p, v[i]*3)
		}
	}

	f.Print(p)
}
