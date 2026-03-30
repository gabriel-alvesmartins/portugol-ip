package main
import "fmt"

func main() {
	var n1, n2, n3, concN int
	fmt.Println("Digite os 3 algorismos do número desejado, separados por ENTER:")
	fmt.Scan(&n1, &n2, &n3)

	if n1 >= 10 || n2 >= 10 || n3 >= 10 {
		fmt.Println("DIGITO INVALIDO")
	} else {
		if n1 != 0 {
			concN = n1*100 + n2*10 + n3
		} else {
			if n2 != 0 {
				concN = n2*10 + n3
			} else {
				concN = n3
				if n3 == 0 {
					concN = 0
				}
			}
		}
		fmt.Println("O seu número é:", concN)
		fmt.Println("O seu número ao quadrado é:", concN*concN)
	}
}
