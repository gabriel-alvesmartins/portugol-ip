package main

import f "fmt"

func main() {

	var v [10]int

	for i := 0; i < 10; i++ {
		f.Scan(&v[i])
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9-i; j++ {
			if v[j] > v[j+1] {
				temp := v[j]       
    			v[j] = v[j+1]  
    			v[j+1] = temp
			}
		}
	}
	f.Print(v)
}
