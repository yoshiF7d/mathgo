package list

type Node struct {
	next,prev *Node
	list *List
	Elem any
}

func (self *Node) Next() *Node {
	if p := self.next; self.list != nil && p != &self.list.root {
		return p
	} 
	return nil
}

func (self *Node) Prev() *Node {
	if p := self.prev; self.list != nil && p != &self.list.root {
		return p
	} 
	return nil
}

type List struct {
	symbol *Symbol
	root *Node
	len int
}

func New() *List {return new(List).Init()}

func (self *List) Init() *List {
	self.root.next = &self.root
	self.root.prev = &self.root
	self.len = 0
	return self
}

func (self *List) Len() int { return self.len }

func (self *List) First() *Node {
	if self.len == 0 {
		return nil
	}
	return self.root.next
}

func (self *List) Last() *Node {
	if self.len == 0 {
		return nil
	}
	return self.root.prev
}

func (self *List) insertNode(node,at *Node) *List {
	node.prev = at
	node.next = at.next
	node.prev.next = node
	node.next.prev = node
	node.list = self
	self.len ++
	return self
}

func (self *List) insert(elem any, at *Node) *List {
	return self.insertNode(&Node{Elem : elem},at)
}

func (self *List) Append(elem any) *List {
	return self.insert(elem,&self.root)
}

func (self *List) Prepend(elem any) *List {
	return self.insert(elem,&self.root.prev)
}