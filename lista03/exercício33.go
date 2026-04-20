package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Scan(&n1,&n2)
	var q int
	var resto int = 0;
	for{
		if n1>=n2{
		n1 = n1-n2
		q++;
		} else{
			resto = n1
			break
		}
	}
	f.Printf("Resto: %d Quociente: %d", resto, q)
}
