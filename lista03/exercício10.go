package main

import f "fmt"

func main() {
	var n int
	f.Scan(&n)
	var a int = 0
	var b int = 1
	f.Print(a, b, " ")
	for i := 2; i < n; i++ {
		var temp int = a + b
		a = b
		b = temp
		f.Print(temp, " ")
	}
}
