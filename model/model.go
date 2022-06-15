package model

type Node struct {
	data string
	next *Node
}

type SingleLinkedList struct {
	size int
	head *Node
}

func (sl *SingleLinkedList) AddToHead(inp string) {
	newNode := &Node{
		data: inp,
	}
	if sl.head == nil {
		sl.head = newNode
	} else {
		newNode.next = sl.head
		sl.head = newNode
		sl.size++
	}
}

func (sl *SingleLinkedList) AddToTail(inp string) {
	newNode := &Node{
		data: inp,
	}
	if sl.head == nil {
		sl.head = newNode
	} else {
		current := sl.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	sl.size++
}
