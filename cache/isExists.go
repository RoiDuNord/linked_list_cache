package cache

import structs "workWithCache/server/convertingResponseToIntSlice/convertingFuncs/intSliceToLinkedList"

func isExists(number int, current *structs.Node) bool {
	for current != nil {
		currentNum := current.Data.Number
		if currentNum == number {
			return true
		}
		current = current.Next
	}
	return false
}
