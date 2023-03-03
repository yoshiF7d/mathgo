package parser

import "fmt"
import "container/list"
import "github.com/yoshiF7d/mathgo/symbol"

var tree list.List

const (
	LIST = iota
	TERMINAL
)

func init() {
	tree.Init()
	for _,s := range symbol.SymbolMap {
		appendToTree(s)
	}
}

func appendToTree(s symbol.SymbolType) {

}
