package cache

import "server/config"

type Cache struct {
	List *LinkedList
}

func New(cfg config.Config) *Cache {
	return &Cache{List: newLinkedList()}
}
