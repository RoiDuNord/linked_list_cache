package cache

func (cache Cache) UpdateWithNewFactor(factor int) {
	cur := cache.List.Head
	for cur != nil {
		cur.Data.MultipliedNumber = cur.Data.Number * factor
		cur = cur.Next
	}
}
