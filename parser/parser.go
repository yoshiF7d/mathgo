package parser
import "fmt"
import "github.com/yoshiF7d/mathgo/list"

func main() {
	l := list.New()
	l.Append("apple")
	l.Append(2)
	l.Append('b')
	l.Prepend(4)

	for e := l.First(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
