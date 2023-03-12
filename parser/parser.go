package parser

import (
	"fmt"
	. "github.com/yoshiF7d/mathgo/symbol"
)
var tree Node

func Parse(s string) *Node{
	return Node_create(&Symbol,nil)
} 

func init() {
	tree.Init()
	for _,s := range SymbolMap {
		if len(s.Format) == 1 {
			appendToTree(s)
		}
	}
	fmt.Println(tree)
}

func appendToTree(s *SymbolType) {
	root := &tree
	for _, cf := range s.Format {
		e := root.Front()
		for ; e != nil; e = e.Next() {
			node := e.Value.(*Node)
			cn := node.Data.(rune)

			if cf == cn {
				root = node
				break
			}
		}
		if e == nil {
			root.PushBack(&Node{Symbol:s,Data:cf})
		}
	}
}
