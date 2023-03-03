package symbol

func init() {
	Comma.Format = `"," .`
	Comma.Precedence = 0
	Comma.Attributes = []string{"Protected"}
}
