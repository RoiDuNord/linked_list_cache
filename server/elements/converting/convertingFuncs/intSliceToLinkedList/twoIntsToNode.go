package convfuncs

type Data struct {
	Number           int
	MultipliedNumber int
}

type Node struct {
	Data Data
	Next *Node
}

func NewNode(number, multipliedNumber int) *Node {
	return &Node{
		Data: Data{
			Number:           number,
			MultipliedNumber: multipliedNumber,
		},
	}
}
