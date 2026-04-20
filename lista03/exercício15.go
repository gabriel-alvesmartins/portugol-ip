package main
import f "fmt"

func main() {
	var n int
	f.Scan(&n)

	for i := 1; i <= n; i++ {
		f.Print(i*i, " ")
	}
}
