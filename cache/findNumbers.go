package cache

import "log"

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
