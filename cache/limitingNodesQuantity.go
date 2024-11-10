package cache

import (
	customerrors "workWithCache/server/customErrors"
)

func (cache *Cache) LimitingNodesQuantity(newSize int) error {
	if newSize < 0 {
		return customerrors.NewNegativeCacheSize(newSize)
	}
	currentNode := cache.Data.Head
	var count int
	for count < newSize {
		currentNode = currentNode.Next
		count++
	}
	currentNode.Next = nil
	return nil
}
