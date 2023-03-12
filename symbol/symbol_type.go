package symbol

import (
	"container/list"
)

type SymbolType struct {
	Name          string
	Format        string
	Precedence    int
	Attributes    []string
	Associativity []string
	ID            uint16
	Evaluate      func(*Node) *Node
}

type Node struct {
	list.List
	Symbol *SymbolType
	Data   any
}

var SymbolMap = map[string]*SymbolType{}

func (s *SymbolType) String() string {
	return s.Name
}

func (n Node) String() string {
	return TreeForm_String(&n)
}

func Node_create(symbol *SymbolType, data any) *Node {
	node := Node{Symbol: symbol, Data: data}
	node.Init()
	return &node
}
