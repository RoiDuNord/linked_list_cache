package convfuncs

import "fmt"

type LinkedList struct {
	Head *Node
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) Append(node Node) {
	if list.Head == nil {
		list.Head = &node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = &node
	}
}

func (list *LinkedList) Output() {
	current := list.Head
	c := 1
	if current == nil {
		return
	}
	for current.Next != nil {
		fmt.Printf("Node №: %d (%d -> %d)\n", c, current.Data.Number, current.Data.MultipliedNumber)
		current = current.Next
		c++
	}
}
