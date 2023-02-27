package symbol

var Comma = symbol{}

func init() {
	Comma.Name = "Comma"
	Comma.Format = `"," .`
	Comma.Precedence = 0
	Comma.Attributes = []string{"Protected"}
}