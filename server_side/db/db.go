// package db

// import (
// 	"errors"
// 	"log"
// 	"time"
// )

// type Database struct {
// 	data []int
// }

// func New() *Database {
// 	return &Database{}
// }

// func (db *Database) FindNumbers(uncachedNumbers []int, factor int) ([]int, error) {
// 	// factor лучше вынести как зависимость БД, чтобы эту проверку не делать и её настройку отдать конструктору
// 	// if factor == 0 {
// 	// 	return nil, errors.New("factor cannot be zero")
// 	// }

// 	if len(uncachedNumbers) == 0 {
// 		return nil, errors.New("all these numbers are already in cache")
// 	}

// 	multipliedNumbers := multiplyMissingNumbers(uncachedNumbers, factor)

// 	db.data = append(db.data, uncachedNumbers...)

// 	log.Printf("All database requests: %d\n", db.data)
// 	return multipliedNumbers, nil
// }

// func multiplyMissingNumbers(missingNumbers []int, factor int) []int {
// 	multipliedNumbers := make([]int, 0, len(missingNumbers))
// 	for _, number := range missingNumbers {
// 		time.Sleep(time.Millisecond * 200)
// 		multipliedNumbers = append(multipliedNumbers, number*factor)
// 	}
// 	return multipliedNumbers
// }

package db

import (
	"log"
	"time"
	"server/config"
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

	log.Printf("all database requests count: %d\n", len(db.Data))
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
