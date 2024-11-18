package cache

import "workWithCache/config"

type Cache struct {
	List *LinkedList
}

func New(cfg config.Config) *Cache {
	return &Cache{List: newLinkedList()}
}
