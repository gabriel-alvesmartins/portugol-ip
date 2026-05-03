package main

import f "fmt"

func main() {

	var n [10]int
	var d[5]int

	for i:=0; i<10; i++{
		f.Scan(&n[i])
	}

	for i:=0; i<5; i++{
		f.Scan(&d[i])
	}

	for i:=0; i<10; i++{
		f.Println("Número: ", n[i])
		for j:=0;j<5;j++{
			if n[i]%d[j]==0{
				f.Printf("Divisível por %d na posição %d\n", d[j], j)
			}
		}
	}
	
}
