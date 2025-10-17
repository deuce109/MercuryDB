package collections

import "sync"

type Collection[T comparable] struct {
	items       []T
	mutex       *sync.Mutex
	mutexLocked bool
}

type MapableCollection[T comparable, R comparable] struct {
	items       []T
	mutex       *sync.Mutex
	mutexLocked bool
}

func MapableCollectionFromCollection[T, R comparable](c *Collection[T]) *MapableCollection[T, R] {
	return &MapableCollection[T, R]{items: c.items, mutex: c.mutex, mutexLocked: c.mutexLocked}
}

func NewCollection[T comparable](items []T) *Collection[T] {
	return &Collection[T]{items: items, mutex: &sync.Mutex{}, mutexLocked: false}
}

func collectionFromMapableCollection[T, R comparable](c *MapableCollection[T, R], newItems []R) *Collection[R] {
	return &Collection[R]{items: newItems, mutex: c.mutex, mutexLocked: c.mutexLocked}
}

func NewCollectionFromString(input string) *Collection[rune] {

	runeSlice := make([]rune, 0, len(input))

	for _, r := range input {
		runeSlice = append(runeSlice, r)
	}

	return &Collection[rune]{items: runeSlice}
}

func NewByteCollectionFromString(input string) *Collection[byte] {

	return &Collection[byte]{items: []byte(input)}
}

func (c *Collection[T]) lock() {
	if !c.mutexLocked {
		c.mutex.Lock()
		c.mutexLocked = true
	}
}

func (c *Collection[T]) unlock() {
	if c.mutexLocked {
		c.mutex.Unlock()
		c.mutexLocked = false
	}
}
