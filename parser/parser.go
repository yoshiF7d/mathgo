package parser

import . "github.com/yoshiF7d/mathgo/symbol"

var tree Node

func init() {
	tree.Init()
	for _,s := range SymbolMap {
		if len(s.Format) == 1 {
			appendToTree(s)
		}
	}
}

func appendToTree(s *SymbolType) {
	root := tree
	for _, cf := range s.Format {
		e := root.Front()
		for ; e != nil; e = e.Next() {
			node := e.Value.(Node)
			cn := node.Data.(rune)

			if cf == cn {
				root = node
				break
			}
		}
		if e == nil {
			root.PushBack(Node{Symbol:SymbolMap["Symbol"],Data:cf})
		}
	}
}
