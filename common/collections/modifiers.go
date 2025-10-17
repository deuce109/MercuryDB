package collections

func (c *Collection[T]) Filter(predicate func(T) bool) *Collection[T] {
	c.lock()
	filteredItems := make([]T, 0, len(c.items))

	for _, item := range c.items {
		if predicate(item) {
			filteredItems = append(filteredItems, item)
		}
	}
	c.items = filteredItems
	return c
}

func (c *Collection[T]) FilterNonZero() *Collection[T] {
	c.lock()
	nonZeroFilterFunc := func(item T) bool {
		var zero T
		return item != zero
	}

	return c.Filter(nonZeroFilterFunc)
}

func (c *MapableCollection[T, R]) Map(mapper func(T) R) *Collection[R] {
	mappedItems := make([]R, 0, len(c.items))
	for _, item := range c.items {
		mappedItems = append(mappedItems, mapper(item))
	}
	return collectionFromMapableCollection(c, mappedItems)
}
