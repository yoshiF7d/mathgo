// Code generated by go generate DO NOT EDIT.
package symbol

const (
	Alternatives_ID = iota
	Bracket_ID      = iota
	Comma_ID        = iota
	Dot_ID          = iota
	DoubleQuotes_ID = iota
	Integer_ID      = iota
	Node_ID         = iota
	Symbol_ID       = iota
)

var Alternatives = SymbolType{}
var Bracket = SymbolType{}
var Comma = SymbolType{}
var Dot = SymbolType{}
var DoubleQuotes = SymbolType{}
var Integer = SymbolType{}
var Node = SymbolType{}
var Symbol = SymbolType{}

func init() {
	Alternatives.ID = Alternatives_ID
	Alternatives.Name = "Alternatives"
	SymbolMap[Alternatives.Name] = Alternatives
	Bracket.ID = Bracket_ID
	Bracket.Name = "Bracket"
	SymbolMap[Bracket.Name] = Bracket
	Comma.ID = Comma_ID
	Comma.Name = "Comma"
	SymbolMap[Comma.Name] = Comma
	Dot.ID = Dot_ID
	Dot.Name = "Dot"
	SymbolMap[Dot.Name] = Dot
	DoubleQuotes.ID = DoubleQuotes_ID
	DoubleQuotes.Name = "DoubleQuotes"
	SymbolMap[DoubleQuotes.Name] = DoubleQuotes
	Integer.ID = Integer_ID
	Integer.Name = "Integer"
	SymbolMap[Integer.Name] = Integer
	Node.ID = Node_ID
	Node.Name = "Node"
	SymbolMap[Node.Name] = Node
	Symbol.ID = Symbol_ID
	Symbol.Name = "Symbol"
	SymbolMap[Symbol.Name] = Symbol
}