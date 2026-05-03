package main

import f "fmt"

func main() {
	var v[15]float64;

	for i:=0; i<15; i++{
		f.Scan(&v[i])
	}

	f.Println("-------------TABELA-------------")
	f.Printf("%-10s %-10s %-10s\n", "NOTA", "F.A", "F.R")

	for i:=0; i<=10; i++{
		count:=0;
		
		for j:=0; j<15; j++{
			if float64(i)==v[j]{
				count++
			}
		}
		freq:=float64(count)/float64(15)*100
		f.Printf("%-10.1d %-10d %-10.2f", i, count, freq)
		f.Print("%\n")
	}
}
