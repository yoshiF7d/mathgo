package parser

import (
	"io"
	"fmt"
	. "github.com/yoshiF7d/mathgo/symbol"
)
var tree *Node

func init() {
	tree = NewNodeSymbol("root")
	for _,s := range SymbolMap {
		if length := len(s.Format); 0 < length && length <= 2 {
			appendToTree(s)
		}
	}
	fmt.Println(TreeForm_String(tree))
}

func appendToTree(s *SymbolType) {
	root := tree
	term := true
	for _, c := range s.Format {
		root,term = lookUpTree(root,c)
		
		if term {
			node := NewNodeSymbol(string(c))
			root.PushBack(node)
			root = node
		}
	}
	root.PushBack(NewNode(s,nil))
}

func lookUpTree(root *Node, c rune) (*Node,bool){
	term := true
	for e := root.Front(); e != nil; e = e.Next() {
		node := e.Value.(*Node)
		cn := rune(node.String()[0])

		if c == cn {
			root = node
			term = false
			break
		}
	}
	
	return root,term
}


func Assign(reader io.RuneReader) *Node{
	root := tree
	term := true

	for {
		if c,_,err:=reader.ReadRune();err!=nil{
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
			}
		}else{
			root,term = lookUpTree(root,c)
			if term {
				break
			}
		}
	}
	return root
}


func Tokenize(str string) *Node{
	stack := NewNodeSymbol("root")
	
	


	return stack
}

func Parse(s string) *Node{
	return NewNode(&Symbol,nil)
} 