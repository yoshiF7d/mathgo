package symbol

var Integer = symbol{}

func init() {
	Integer.Name = "Integer"
	Integer.Format = 
`
integer = decimel {decimal} .
decimal = "0" ... "9" .
`
}