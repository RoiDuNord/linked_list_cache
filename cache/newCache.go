package cache

import "workWithCache/config"

type Cache struct {
	Data *LinkedList
}

func New(cfg config.Config) *Cache {
	return &Cache{Data: newLinkedList()}
}
