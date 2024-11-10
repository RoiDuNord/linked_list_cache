package db

import (
	"errors"
	"log"
	"net/http" // ну не должна у тебя сетевая зависимость быть здесь
	"time"
	"workWithCache/cache" // и пакет кэша тоже
)

// Слой базы данных должен зависеть только от драйера БД (здесь не должно быть)
// Работу с кэшом следует вынести на уровень выше по логике

type Database struct {
	data []int
}

func New() *Database {
	return &Database{}
}

// по логике в целом всё отлично
// если уж прям сильно понадобилось иметь связь с кэшом из БД, то этот несчастный указатель надо один раз положить в в саму базу, а не таскать к каждой функции
func (db *Database) FindNumbers(w http.ResponseWriter, inputSlice []int, factor int, cache *cache.Cache) ([]int, error) {
	// factor лучше вынести как зависмость БД, чтобы эту проверку не делать и её настройку отдать конструктору
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

// функция это глагол - HandleMissingNumbers
func missingNumbersHandle(missingNumbers []int, factor int, cache *cache.Cache) []int {
	multipliedNumbers := make([]int, 0, len(missingNumbers))
	for _, number := range missingNumbers {
		time.Sleep(time.Millisecond * 200)
		multipliedNumbers = append(multipliedNumbers, number*factor)
		cache.InsertNumber(number, number*factor)
	}
	return multipliedNumbers
}
