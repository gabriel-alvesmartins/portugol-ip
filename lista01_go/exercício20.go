package main
import "fmt"

func main() {
	var h, m, s, ts float64

	fmt.Println("Digite o valor em horas, minutos e segundos, respectivamente, para saber o total em segundos")
	fmt.Scan(&h, &m, &s)

	ts = h*3600 + m*60 + s
	fmt.Println("O TEMPO EM SEGUNDOS É:", ts)
}
