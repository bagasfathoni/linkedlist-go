package main

import (
	"fmt"
	"linked-list/model"
)

func main() {
	singleList := InitNewLinkedList()
	/*	Create a new linked list :
		++++	++++	++++	++++
		+25+ -- +15+ -- +80+ -- +75+
		++++	++++	++++	++++
	*/
	singleList.AddToHead("80")
	singleList.AddToHead("15")
	singleList.AddToHead("25")
	singleList.AddToTail("75")

	/*	Then add:
		++++	++++	++++	++++	++++
		+25+ -- +15+ -- +80+ -- +35+ -- +75+
		++++	++++	++++	++++	++++
	*/
	newList := model.AddAfterNode(*singleList, "15", "35")

	// To view value of a node edit here.
	/* Example :
	First node --> list.Head.Data
	Second node --> list.Head.Next.Data
	Third node --> list.Head.Next.NextData
	*/
	fmt.Println(newList.Head.Next.Next.Data)

}

func InitNewLinkedList() *model.SingleLinkedList { // Initialize an empty linked list
	return &model.SingleLinkedList{}
}
