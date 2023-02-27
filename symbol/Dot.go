package symbol

var Dot = symbol{}

func init() {
	Dot.Name = "Dot"
	Dot.Format = `"." .`
	Dot.Precedence = 490
	Dot.Attributes = []string{"Protected"}
}