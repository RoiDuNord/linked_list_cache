package cache

type Cache struct {
	Data *LinkedList
}

func New() *Cache {
	return &Cache{Data: NewLinkedList()}
}
