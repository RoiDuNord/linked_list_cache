package cache

func isExists(number int, current *Node) bool {
	for current != nil {
		currentNum := current.data.number
		if currentNum == number {
			return true
		}
		current = current.next
	}
	return false
}
