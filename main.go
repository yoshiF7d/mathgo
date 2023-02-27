package main
import "fmt"
//import "github.com/yoshiF7d/mathgo/list"
import "github.com/yoshiF7d/mathgo/symbol"

func main() {
	s := symbol.Integer
	fmt.Println(s.Name)
	fmt.Println(s.Attributes)
	fmt.Println(s.Format)
}
