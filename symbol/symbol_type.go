package symbol

import (
	"container/list"
	"bytes"
)

type SymbolType struct {
	Name          string
	Format        string
	Precedence    int
	Attributes    []string
	Associativity []string
	ID            uint16
}

type Node struct {
	list.List
	Symbol *SymbolType
	Data   any
}

var SymbolMap = map[string]*SymbolType{}

func (s *SymbolType) String() string{
	return s.Name
}

func (n *Node) String() string{

}

func (n *Node) StringMod(level int) string{
	var buf bytes.Buffer
	for i:=0 ; i<level; i++ {
		buf.WriteString("\t")
	}
}