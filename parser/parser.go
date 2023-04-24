package parser

import (
	"io"
	"log"
	"strings"
	"github.com/yoshiF7d/mathgo/symbol"
)

var tree *symbol.Node

func init() {
	tree = symbol.NewNode(&symbol.Symbol,"root")
	for _,s := range symbol.SymbolList {
		if length := len(s.Format); 0 < length && length <= 3 {
			appendToTree(s)
		}
	}
	//fmt.Println(TreeForm_String(tree))
}

func appendToTree(s *symbol.SymbolType) {
	root,hit := tree,false
	for _, c := range s.Format {
		root,hit = lookUpTree(root,c)
		if !hit {
			node := symbol.NewNode(&symbol.Symbol,string(c))
			root.Append(node)
			root = node
		}
	}
	root.Append(symbol.NewNode(s,nil))
}

func lookUpTree(root *symbol.Node, c rune) (*symbol.Node,bool){
	hit := false
	for e := root.First(); e != nil; e = e.Next() {
		node := e.Value
		cn := rune(node.String()[0])

		//fmt.Println(node.String())
		
		if c == cn {
			root = node
			hit = true
			break
		}
	}
	return root,hit
}

func assign(reader *strings.Reader) (*symbol.Node,error){
	root,hit := tree,false
	//pare := tree
	var (
		c rune
		err error
	)
	for {
		if c,_,err = reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		} else {
			//pare = root
			root,hit = lookUpTree(root,c)
			//fmt.Println(string(c))
			if hit {
				if root.Len() == 0 {
					break
				}
			} else {
				reader.UnreadRune()
				break
			}
		}
	}
	if root != tree {
		for e:= root.First(); e != nil; e = e.Next() {
			node := e.Value
			if node.Symbol.ID != symbol.Symbol_ID {
				root = node
				break
			}
		}
	}
	return root,err
}

func Tokenize(s string) *symbol.Node{
	stack := symbol.NewNode(&symbol.Symbol,"root")
	reader := strings.NewReader(s)
	for {
		node,err := assign(reader)
		if node != tree {
			stack.Append(node)	
		}
		
		if node == tree || err == io.EOF{
			break
		}
	}
	return stack
}

func Parse(s string) *symbol.Node{
	return Tokenize(s)
} 

func Root() *symbol.Node{
	return tree
}