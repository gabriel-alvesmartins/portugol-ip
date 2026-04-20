package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Scan(&n1,&n2)

	a:=n1
	b:=n2
	
	for b!=0{
		temp:=b
		b=a%b
		a=temp
	}
	mmc := (n1 / a) * n2
	f.Println(mmc)
}
