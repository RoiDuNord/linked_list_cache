package cache

import structs "workWithCache/server/convertingResponseToIntSlice/convertingFuncs/intSliceToLinkedList"

func (cache *Cache) InsertNumber(missingNumber, multipliedMissingNumber int) {
	node := structs.NewNode(missingNumber, multipliedMissingNumber)
	cache.Data.Append(*node)
}
