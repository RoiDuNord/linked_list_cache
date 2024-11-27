package cache

import (
	"log"
	customerrors "server/server/customErrors"
)

func (cache *Cache) LimitingNodesQuantity(newSize int) error {
	if newSize < 0 {
		return customerrors.NegativeCacheSizeError(newSize)
	}

	currentNode := cache.List.Head
	var count int

	for currentNode != nil {
		count++
		currentNode = currentNode.Next
	}

	if count < newSize {
		return customerrors.ExceedingCacheSizeError(count, newSize)
	}

	if count == newSize {
		log.Printf("it is already your cache size (%d)\n", newSize)
	}

	currentNode = cache.List.Head
	count = 0

	for count < newSize-1 {
		currentNode = currentNode.Next
		count++
	}

	if currentNode != nil {
		currentNode.Next = nil
	}

	return nil
}
