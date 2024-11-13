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

	for currentNode != nil {
		count++
		currentNode = currentNode.next
	}

	if count <= newSize {
		return nil
	}

	currentNode = cache.Data.head
	count = 0

	for count < newSize-1 {
		currentNode = currentNode.next
		count++
	}

	if currentNode != nil {
		currentNode.next = nil
	}

	return nil
}
