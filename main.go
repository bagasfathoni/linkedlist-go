package main

import (
	"linked-list/model"
)

func main() {
	singleList := InitNewLinkedList()
	singleList.AddToHead("80")
	singleList.AddToHead("15")
	// singleList.AddToHead("15")
	singleList.AddToTail("15")
	// fmt.Println(singleList.Head.Next.Data)

}

func InitNewLinkedList() *model.SingleLinkedList {
	return &model.SingleLinkedList{}
}
