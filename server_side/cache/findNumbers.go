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
					break
				}
				current = current.Next
			}
		} else {
			missingNumbers = append(missingNumbers, number)
		}
	}

	log.Printf("values from found numbers: %d", multipliedNumbers)
	log.Printf("missing numbers: %d", missingNumbers)
	return multipliedNumbers, missingNumbers
}
