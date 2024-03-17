package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func NewLinkList() *LinkedList {
	return &LinkedList{head: &Node{}}
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
		return
	}

	newNode.next = list.head
	list.head = newNode
}

func (list *LinkedList) Display() {
	if list.head == nil {
		return
	}

	current := list.head
	for current != nil {
		fmt.Printf("Val: %d -> ", current.data)
		current = current.next
	}
}
