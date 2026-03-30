package main
import "fmt"

func main() {
	var salario, energia, custo, porKw, desconto float64

	fmt.Print("Qual o valor do salário mínimo? ")
	fmt.Scan(&salario)
	fmt.Print("Quanto que é gasto de energia na sua casa(em kW)? ")
	fmt.Scan(&energia)

	custo = energia * 0.007 * salario
	porKw = custo / energia
	desconto = custo * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", porKw)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custo)
	fmt.Printf("Custo com desconto: R$ %.2f\n", desconto)
}
