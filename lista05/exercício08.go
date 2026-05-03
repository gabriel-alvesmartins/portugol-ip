package main

import f "fmt"
import "math"

func main() {
	var v[]float64;

	for i:=0; i<15; i++{
		var a float64
		f.Scan(&a)

		if a<0{
			v = append(v, -1)
		}else{
			b:= math.Sqrt(a)
			v = append(v, b)
		}
	}

	f.Print(v)
	
	
}
