package main

import f "fmt"

func main() {
	var n1, n2 int
	f.Scan(&n1, &n2)
	for i := n1; i <= n2; i++ {
		var p int = 0;
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				p++;			
			}
			if j==i && p==2{
				f.Println(i)
			}
		}
	}
}
