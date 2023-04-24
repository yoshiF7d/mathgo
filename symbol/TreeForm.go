package symbol

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
)

func init() {
	TreeForm.Evaluate = TreeForm_Evaluate
}

func TreeForm_Evaluate(node *Node) *Node {
	return NewNode(&String, TreeForm_String(node))
}

func TreeForm_print(node *Node) {
	fmt.Println(TreeForm_String(node))
}

func TreeForm_String(node *Node) string {
	buf := strings.Builder{}
	TreeForm_StringMod(&buf, node, []byte{})
	return buf.String()
}

func TreeForm_StringMod(buf *strings.Builder, node *Node, indent []byte) {
	var name string

	if node.Symbol == nil {
		name = ""
	} else {
		name = node.String()
	}

	buf.WriteString(color.HiBlueString(name))
	blen := len(indent)
	nlen := len(name)

	ind := make([]byte, nlen+blen+3)
	for i := 0; i < nlen+blen; i++ {
		ind[i] = ' '
	}
	if len := node.Len(); len > 0 {
		if len > 1 {
			ind[nlen+blen] = ' '
			ind[nlen+blen+1] = '|'
			ind[nlen+blen+2] = ' '
			buf.WriteString("-+-")
			TreeForm_StringMod(buf, node.First().Value, ind)
			for e := node.First().Next(); e.Next() != nil; e = e.Next() {
				buf.Write(indent)
				for i := 0; i < nlen; i++ {
					buf.WriteByte(' ')
				}
				buf.WriteString(" +-")
				TreeForm_StringMod(buf, e.Value, ind)
			}
			ind[nlen+blen] = ' '
			ind[nlen+blen+1] = ' '
			ind[nlen+blen+2] = ' '
			buf.Write(indent)
			for i := 0; i < nlen; i++ {
				buf.WriteByte(' ')
			}
			buf.WriteString(" \\-")
			TreeForm_StringMod(buf, node.Last().Value, ind)
		} else {
			ind[nlen+blen] = ' '
			ind[nlen+blen+1] = ' '
			ind[nlen+blen+2] = ' '
			buf.WriteString("---")
			TreeForm_StringMod(buf, node.Last().Value, ind)
		}
	} else {
		buf.WriteByte('\n')
	}
}
