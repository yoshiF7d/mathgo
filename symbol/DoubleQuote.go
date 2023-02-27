package symbol

var DoubleQuotes = symbol{}

func init() {
	DoubleQuotes.Name = "Dot"
	DoubleQuotes.Format = `"." .`
	DoubleQuotes.Precedence = 490
	DoubleQuotes.Attributes = []string{"Protected"}
}