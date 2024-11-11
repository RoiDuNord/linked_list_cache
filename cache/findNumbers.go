package cache

import "log"

func (cache *Cache) FindNumbers(inputSlice []int) ([]int, []int) {
	var multipliedNumbers, missingNumbers []int
	list := cache.Data

	for _, number := range inputSlice {
		current := list.head
		if isExists(number, current) {
			for current != nil {
				if current.data.number == number {
					multipliedNumber := current.data.multipliedNumber
					multipliedNumbers = append(multipliedNumbers, multipliedNumber)
					log.Printf("%d exists in cache with multiplied value %d\n", number, current.data.multipliedNumber)
					break
				}
				current = current.next
			}
		} else {
			missingNumbers = append(missingNumbers, number)
			log.Printf("%d is not in cache\n", number)
		}
	}

	log.Printf("Values from found numbers: %v", multipliedNumbers)
	log.Printf("Missing numbers: %v", missingNumbers)
	return multipliedNumbers, missingNumbers
}
