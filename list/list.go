package list

type Node struct {
	next,prev *Node
	list *List
	Value any
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
	root Node
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

func (self *List) insert(value any, at *Node) *List {
	return self.insertNode(&Node{Value : value},at)
}

func (self *List) Prepend(value any) *List {
	return self.insert(value,&self.root)
}

func (self *List) Append(value any) *List {
	return self.insert(value,self.root.prev)
}
