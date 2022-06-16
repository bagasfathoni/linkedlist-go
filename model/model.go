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

func (sl *SingleLinkedList) checkSameValue(inp *Node) error { // Linked list can not have a same value, use this method to apply the rule

	current := sl.Head
	for current.Next != nil {
		if current.Data == inp.Data {
			return errors.New("can not enter a same value: ")
		}
		current = current.Next
	}
	return nil
}

func AddAfterNode(sl SingleLinkedList, after, input string) SingleLinkedList {
	var newLinkedList SingleLinkedList
	// var new Node
	inputNode := &Node{
		Data: input,
	}
	current := sl.Head
	if current.Next == nil { // Linked list only have 1 value
		current.Next = inputNode
	}
	if current.Next != nil { // Linked list have more than  1 value
		for current.Data != after {
			newLinkedList.AddToTail(current.Data)
			current = current.Next
		}
		newLinkedList.AddToTail(current.Data)
		newLinkedList.AddToTail(input)
		for current.Next != nil {
			newLinkedList.AddToTail(current.Next.Data)
			current = current.Next
		}
	}
	return newLinkedList

}
func (sl *SingleLinkedList) FindNode(val string) {
	current := sl.Head
	if current.Data != val {
		for current.Data != val {
			current = current.Next
		}
		fmt.Println("there is a matched value")
	} else if current.Data == val {
		fmt.Println("there is a matched value")
	} else {
		fmt.Println("no matched value")
	}
}
