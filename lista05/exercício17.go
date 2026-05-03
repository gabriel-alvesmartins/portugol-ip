package main

import f "fmt"

func main() {

	var v [10]int
	
	for i:=0; i<10; i++{
		f.Scan(&v[i])
		count:=0
		for z:=1; z<=v[i];z++{
			
			if v[i]%z==0{
				count++
			}
		}

		if count==2{
				f.Printf("O número %d é primo. Ocupa a posição %d\n", v[i], i)
			}		
	}
	
	
}
