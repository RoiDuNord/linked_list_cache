package db

import (
	"log"
	"server/config"
	"time"
)

type Database struct {
	Data []int
}

func New(cfg config.Config) *Database {
	return &Database{}
}

func (db *Database) FindNumbers(uncachedNumbers []int, factor int) []int {

	multipliedNumbers := multiplyMissingNumbers(uncachedNumbers, factor)

	db.Data = append(db.Data, uncachedNumbers...)

	log.Printf("different requested numbers quantity: %d\n", len(db.Data))
	return multipliedNumbers
}

func multiplyMissingNumbers(missingNumbers []int, factor int) []int {
	multipliedNumbers := make([]int, 0, len(missingNumbers))
	for _, number := range missingNumbers {
		time.Sleep(time.Millisecond * 200)
		multipliedNumbers = append(multipliedNumbers, number*factor)
	}
	return multipliedNumbers
}
