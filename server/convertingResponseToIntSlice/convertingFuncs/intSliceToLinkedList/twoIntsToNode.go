package intslicetolinkedlist

type Data struct {
	Number           int
	MultipliedNumber int
}

type Node struct {
	Data Data
	Next *Node
}

// В вот тут не нравится
// Структура по сути принаделжит кэшу, поэтому класть её в отдельный пакет смысла нет
func NewNode(number, multipliedNumber int) *Node {
	return &Node{
		Data: Data{
			Number:           number,
			MultipliedNumber: multipliedNumber,
		},
	}
}
