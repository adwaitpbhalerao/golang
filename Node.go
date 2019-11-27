package main


// Node struct
type Node struct {
	value interface{}
	next  *Node
	previous *Node
}

//return the current node value
func (n *Node) Value() interface{} {
	return n.value
}

//return the next node if exists
func (n *Node) Next() *Node {
	return n.next
}


// set the next node
func (n *Node) SetNext(next *Node) {
	n.next = next
}

// stack node functions
// return previous node
func (n *Node) Previous() *Node{
	return n.previous
}

// function to set the previous node
func (n *Node) SetPrevious(previous *Node){
	n.previous = previous
}