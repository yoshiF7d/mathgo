package symbol

type symbol struct{
	Name string
	Format string
	Precedence int
	Attributes []string
	Associativity []string
}

var Symbol = symbol{}

func init(){
	Symbol.Name = "Symbol"
	Symbol.Precedence = 1000
}
