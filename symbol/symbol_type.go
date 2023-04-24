package symbol

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

var SymbolMap = map[string]*SymbolType{}
var SymbolList []*SymbolType

func (s *SymbolType) String() string {
	return s.Name
}
