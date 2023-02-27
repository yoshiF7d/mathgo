package symbol

var Bracket = symbol{}

func init() {
	Bracket.Name = "Bracket"
	Bracket.Format = `"[" | "]" .`
	Bracket.Precedence = 1000
	Bracket.Associativity = []string{"Bracket"}
}