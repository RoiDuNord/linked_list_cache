package cache

import (
	customerrors "workWithCache/server/customErrors"
)

func (cache *Cache) LimitingNodesQuantity(newSize int) error {
	if newSize < 0 {
		return customerrors.NewNegativeCacheSize(newSize)
	}
	currentNode := cache.Data.head
	var count int
	for count < newSize {
		currentNode = currentNode.next
		count++
	}
	currentNode.next = nil
	return nil
}
