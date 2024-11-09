package cache

// здесь мапа кеш сырки мешки похуй

import (
	"log"
	structs "workWithCache/server/elements/converting/convertingFuncs/intSliceToLinkedList"
)

type Cache struct {
	Data *structs.LinkedList
}

func New() *Cache {
	return &Cache{Data: structs.NewLinkedList()}
}

func (cache *Cache) FindNumbers(inputSlice []int) ([]int, []int) {
	var foundNumbers, missingNumbers []int
	list := cache.Data

	for _, number := range inputSlice {
		current := list.Head
		if isExists(number, current) {
			foundNumbers = append(foundNumbers, number)
			log.Printf("%d exists in cache\n", number)
		} else {
			missingNumbers = append(missingNumbers, number)
			log.Printf("%d is not in cache\n", number)
		}
	}

	log.Printf("Values from found numbers: %v", foundNumbers)
	log.Printf("Missing numbers: %v", missingNumbers)
	return foundNumbers, missingNumbers
}

func (cache *Cache) InsertNumber(missingNumber, multipliedMissingNumber int) {
	node := structs.NewNode(missingNumber, multipliedMissingNumber)
	cache.Data.Append(*node)
}

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
