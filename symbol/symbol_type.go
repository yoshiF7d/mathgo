package symbol

type SymbolType struct {
	Name          string
	Format        string
	Precedence    int
	Attributes    []string
	Associativity []string
	ID            uint16
}

var SymbolMap = map[string]SymbolType{}
