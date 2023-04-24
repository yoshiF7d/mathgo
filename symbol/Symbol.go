package symbol

func init() {
	Symbol.Precedence = 1000
	Symbol.toString = Symbol_toString
	Symbol.Format = `
		integer = letter { letter | unicode_digit } .
		decimal = "0" ... "9" .
	`
}

func Symbol_toString(node *Node) string {
	if node.Data != nil {
		return node.Data.(string)
	} else {
		return ""
	}
}
