package symbol

func init() {
	Symbol.Precedence = 1000
	Symbol.toString = Symbol_toString
}

func Symbol_toString(node *Node) string {
	if node.Data != nil {
		return node.Data.(string)
	} else {
		return ""
	}
}

func NewNodeSymbol(name string) *Node {
	return &Node{Symbol: &Symbol, Data: name}
}
