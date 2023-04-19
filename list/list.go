package list
import "github.com/yoshiF7d/mathgo/symbol"

type Element struct {
	Value *Node
	next,prev *Element
	list *Node
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	} 
	return nil
}

type Node struct {
	Symbol *symbol.SymbolType
	Data   any
	root Element
	len int
}

func New() *Node {return new(Node).Init()}

func (n *Node) Init() *Node {
	n.root.next = &n.root
	n.root.prev = &n.root
	n.len = 0
	return n
}

func (n *Node) Len() int { return n.len }

func (n *Node) First() *Element {
	if n.len == 0 {
		return nil
	}
	return n.root.next
}

func (n *Node) Last() *Element {
	if n.len == 0 {
		return nil
	}
	return n.root.prev
}

func (n *Node) insertNode(e,at *Element) *Node {
	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
	e.list = n
	n.len ++
	return n
}

func (n *Node) insert(n2 *Node, at *Element) *Node {
	return n.insertNode(&Element{Value : n2},at)
}

func (n *Node) Prepend(n2 *Node) *Node {
	return n.insert(n2,&n.root)
}

func (n *Node) Append(n2 *Node) *Node {
	return n.insert(n2,n.root.prev)
}
