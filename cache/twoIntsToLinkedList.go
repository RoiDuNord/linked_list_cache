package cache

import (
	"fmt"
	"net/http"
)

type Data struct {
	number           int
	multipliedNumber int
}

type Node struct {
	data Data
	next *Node
}

type LinkedList struct {
	head *Node
}

func newNode(number, multipliedNumber int) *Node {
	return &Node{
		data: Data{
			number:           number,
			multipliedNumber: multipliedNumber,
		},
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Append(node Node) {
	if list.head == nil {
		list.head = &node
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = &node
	}
}

func (list *LinkedList) Output(w http.ResponseWriter) {
	current := list.head
	c := 1
	if current == nil {
		return
	}
	for current.next != nil {
		fmt.Fprintf(w, "Node №: %d (%d -> %d)\n", c, current.data.number, current.data.multipliedNumber)
		current = current.next
		c++
	}
}
