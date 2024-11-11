package cache

func (cache *Cache) InsertNumber(missingNumber, multipliedMissingNumber int) {
	node := newNode(missingNumber, multipliedMissingNumber)
	cache.Data.append(node)
}
