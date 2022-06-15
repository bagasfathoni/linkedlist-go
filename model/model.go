package model

import (
	"errors"
	"fmt"
)

type Node struct {
	Data string
	Next *Node
}

type SingleLinkedList struct {
	Size int
	Head *Node
}

func (sl *SingleLinkedList) AddToHead(inp string) {
	newNode := &Node{
		Data: inp,
	}
	if sl.Head == nil {
		sl.Head = newNode
	} else {
		err := sl.checkSameValue(newNode)
		if err != nil {
			fmt.Println(err.Error(), newNode.Data)
			return
		}
		newNode.Next = sl.Head
		sl.Head = newNode
		sl.Size++
	}
}

func (sl *SingleLinkedList) AddToTail(inp string) {
	newNode := &Node{
		Data: inp,
	}
	if sl.Head == nil {
		sl.Head = newNode
	} else {
		err := sl.checkSameValue(newNode)
		if err != nil {
			fmt.Println(err.Error(), newNode.Data)
			return
		}
		current := sl.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	sl.Size++
}

func (sl *SingleLinkedList) checkSameValue(inp *Node) error {
	current := sl.Head
	for current.Next != nil {
		if current.Data == inp.Data {
			return errors.New("can not enter a same value: ")
		}
		current = current.Next
	}
	return nil
}
