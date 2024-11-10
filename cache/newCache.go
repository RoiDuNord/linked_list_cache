package cache

import (
	structs "workWithCache/server/convertingResponseToIntSlice/convertingFuncs/intSliceToLinkedList"
)

type Cache struct {
	Data *structs.LinkedList
}

func New() *Cache {
	return &Cache{Data: structs.NewLinkedList()}
}
