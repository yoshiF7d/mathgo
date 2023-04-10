package main

import (
	"fmt"
	"flag"
	"github.com/yoshiF7d/mathgo/parser"
	"github.com/yoshiF7d/mathgo/symbol"
)

func main() {
	flag.Parse()
	args := flag.Args()
	p := parser.Parse(args[0])
	symbol.TreeForm_print(parser.Root())
	fmt.Println(symbol.TreeForm_String(p))
}
