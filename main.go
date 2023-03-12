package main

import (
	"fmt"
	"github.com/yoshiF7d/mathgo/symbol"
	"github.com/yoshiF7d/mathgo/parser"
)

func main() {
	m := symbol.TreeForm
	parser.Parse("test")
	fmt.Println(m)
}
