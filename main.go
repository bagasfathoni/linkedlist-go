package main

import "linked-list/model"

func main() {
	singleList := InitNewLinkedList()
	singleList.AddToHead("80")
	singleList.AddToHead("15")
	singleList.AddToTail("21")
	// fmt.Println(singleList.head.next.next.data)

}

func InitNewLinkedList() *model.SingleLinkedList {
	return &model.SingleLinkedList{}
}
