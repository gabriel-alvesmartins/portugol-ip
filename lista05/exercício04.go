package main

import f "fmt"

func main() {
	var v[10]int;
	var cont[10]bool;
	
	for i:= 0; i<10; i++{
		f.Scan(&v[i])
	}

	for i:=0; i<10; i++{
		if cont[i]{
			continue
		}
		count:=1;
		for j:=i+1; j<10; j++{
			if v[i]==v[j]{
				count++;
				cont[j] = true;
			}
		}
		if count>1 {
			f.Print(v[i], ": ", count, "\n")
		}
	}
}
