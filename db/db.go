package db

import (
	"errors"
	"log"
	"net/http"
	"time"
	"workWithCache/cache"
)

type Database struct {
	data []int
}

func New() *Database {
	return &Database{}
}

func (db *Database) FindNumbers(w http.ResponseWriter, inputSlice []int, factor int, cache *cache.Cache) ([]int, error) {
	if factor == 0 {
		return nil, errors.New("factor cannot be zero")
	}

	cachedNumbers, missingNumbers := cache.FindNumbers(inputSlice)
	if len(missingNumbers) == 0 {
		return cachedNumbers, errors.New("all these numbers are already in cache")
	}

	multipliedNumbers := missingNumbersHandle(missingNumbers, factor, cache)

	multipliedNumbers = append(multipliedNumbers, cachedNumbers...)
	db.data = append(db.data, multipliedNumbers...)

	log.Printf("All database requests: %d\n", db.data)
	return multipliedNumbers, nil
}

func missingNumbersHandle(missingNumbers []int, factor int, cache *cache.Cache) []int {
	multipliedNumbers := make([]int, 0, len(missingNumbers))
	for _, number := range missingNumbers {
		time.Sleep(time.Millisecond * 200)
		multipliedNumbers = append(multipliedNumbers, number*factor)
		cache.InsertNumber(number, number*factor)
	}
	return multipliedNumbers
}
