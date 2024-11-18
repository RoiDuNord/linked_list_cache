package cache

import "log"

func (cache *Cache) FindNumbers(inputSlice []int) ([]int, []int) {
	var multipliedNumbers, missingNumbers []int
	list := cache.List

	for _, number := range inputSlice {
		current := list.Head
		if isExists(number, current) {
			for current != nil {
				if current.Data.Number == number {
					multipliedNumber := current.Data.MultipliedNumber
					multipliedNumbers = append(multipliedNumbers, multipliedNumber)
					log.Printf("%d exists in cache with multiplied value %d\n", number, current.Data.MultipliedNumber)
					break
				}
				current = current.Next
			}
		} else {
			missingNumbers = append(missingNumbers, number)
			log.Printf("%d is not in cache\n", number)
		}
	}

	log.Printf("Values from found numbers: %d", multipliedNumbers)
	log.Printf("Missing numbers: %d", missingNumbers)
	return multipliedNumbers, missingNumbers
}
