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
	toString      func(*Node) string
}

type Node struct {
	list.List
	Symbol *SymbolType
	Data   any
}

var SymbolMap = map[string]*SymbolType{}
var SymbolList []*SymbolType

func (s *SymbolType) String() string {
	return s.Name
}

func (node *Node) String() string {
	if node.Symbol == nil {
		return ""
	} else {
		if node.Symbol.toString == nil {
			return node.Symbol.Name
		} else {
			return node.Symbol.toString(node)
		}
	}
}

func (node *Node) Child() *Node {
	return node.Front().Value.(*Node)
}

func GetNode(e *list.Element) *Node {
	return e.Value.(*Node)
}

func NewNode(symbol *SymbolType, data any) *Node {
	node := Node{Symbol: symbol, Data: data}
	node.Init()
	return &node
}
