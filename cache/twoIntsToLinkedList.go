package cache

import (
	"encoding/json"
	"net/http"
)

type Data struct {
	Number           int `json:"number"`
	MultipliedNumber int `json:"multipliedNumber"`
}

type Node struct {
	Data Data
	Next *Node
}

type LinkedList struct {
	Head *Node
}

func newNode(number, multipliedNumber int) *Node {
	return &Node{
		Data: Data{
			Number:           number,
			MultipliedNumber: multipliedNumber,
		},
	}
}

func newLinkedList() *LinkedList {
	return &LinkedList{}
}

func (list *LinkedList) append(node *Node) {
	if list.Head == nil {
		list.Head = node
	} else {
		current := list.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	}
}

func (list *LinkedList) Output(w http.ResponseWriter) {
	var listNodes CacheNodes
	current := list.Head

	if current == nil {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(listNodes)
		return
	}

	for current != nil {
		listNodes.Nodes = append(listNodes.Nodes, current.Data)
		current = current.Next
	}

	if err := json.NewEncoder(w).Encode(listNodes); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
	}
}

type CacheNodes struct {
	Nodes []Data `json:"list"`
}
