package main

import f "fmt"

func main() {
	var n int
	f.Scan(&n)
	
	if n == 0 {
		f.Println("0")
		return
	}

	var hex []string

	for n > 0 {
		resto := n % 16

		if resto < 10 {
			hex = append(hex, string(resto+'0'))
		} else {
			hex = append(hex, string((resto-10)+'A'))
		}
		n = n / 16
	}

	for i := len(hex) - 1; i >= 0; i-- {
		f.Print(hex[i])
	}
}
