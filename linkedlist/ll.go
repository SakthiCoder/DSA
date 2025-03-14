package linkedlist

type Node struct {
	Previous *Node
	Data     interface{}
	Next     *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (d *DoublyLinkedList) AddAtEnd(val interface{}) {
	newNode := &Node{Data: val}

	if d.Head == nil {
		d.Head = newNode
		d.Tail = newNode
		return
	}

	d.Tail.Next = newNode
	newNode.Previous = d.Tail

	d.Tail = newNode
}
