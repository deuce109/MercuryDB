package collections

func (c *Collection[T]) ToSlice() []T {
	defer c.unlock()
	return c.items
}

func (c *Collection[T]) Count() int {
	defer c.unlock()
	return len(c.items)
}

func (c *Collection[T]) ForEach(action func(T)) {
	defer c.unlock()

	for _, item := range c.items {
		action(item)
	}
}

func (c *Collection[T]) Any(predicate func(T) bool) bool {
	defer c.unlock()

	for _, item := range c.items {
		if predicate(item) {
			return true
		}
	}
	return false
}

func (c *Collection[T]) All(predicate func(T) bool) bool {
	defer c.unlock()

	for _, item := range c.items {
		if !predicate(item) {
			return false
		}
	}
	return true
}
func (c *Collection[T]) First() (T, bool) {
	defer c.unlock()

	var item T
	var zero T
	if len(c.items) > 0 {
		item = c.items[0]
	}

	return c.items[0], item != zero
}

func (c *Collection[T]) Last() (T, bool) {
	defer c.unlock()

	var item T
	var zero T
	if len(c.items) > 0 {
		item = c.items[0]
	}

	return c.items[len(c.items)-1], item != zero
}

func (c *Collection[T]) FirstMatch(predicate func(T) bool) (T, bool) {
	defer c.unlock()

	var result T
	var zero T

	for _, item := range c.items {
		if predicate(item) {
			result = item
			break
		}
	}
	return result, result == zero
}

func (c *Collection[T]) LastMatch(predicate func(T) bool) (T, bool) {

	var result T
	var zero T
	for i := len(c.items) - 1; i >= 0; i-- {
		if predicate(c.items[i]) {
			result = c.items[i]
			break
		}
	}
	defer c.unlock()
	return result, result == zero
}

func (c *Collection[T]) Contains(item T) bool {
	defer c.unlock()
	for _, v := range c.items {
		if v == item {
			return true
		}
	}
	return false
}

func (c *Collection[T]) ContainsAll(items ...T) bool {
	defer c.unlock()
	for _, item := range items {
		if !c.Contains(item) {
			return false
		}
	}

	return true
}

func (c *Collection[T]) Reduce(reducer func(any, T) any, initial any) any {
	defer c.unlock()
	result := initial
	for _, item := range c.items {
		result = reducer(result, item)
	}
	return result
}
