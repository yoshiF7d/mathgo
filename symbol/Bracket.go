package symbol

func init() {
	Bracket.Format = `"[" | "]" .`
	Bracket.Precedence = 1000
	Bracket.Associativity = []string{"Bracket"}
}
