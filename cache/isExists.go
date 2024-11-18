package cache

func isExists(number int, current *Node) bool {
	for current != nil {
		currentNum := current.Data.Number
		if currentNum == number {
			return true
		}
		current = current.Next
	}
	return false
}
