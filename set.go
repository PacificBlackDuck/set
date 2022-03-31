package set

import (
	"sync"
)

type Set[T comparable] struct {
	items map[T]bool
	count uint
	lock  sync.RWMutex
}

func New[T comparable](items ...T) *Set[T] {
	s := &Set[T]{
		items: map[T]bool{},
		count: 0,
		lock:  sync.RWMutex{},
	}

	for _, item := range items {
		s.Add(item)
	}
	return s
}

func NewFromSlice[T comparable](slice []T) *Set[T] {
	s := New[T]()
	for _, item := range slice {
		s.Add(item)
	}
	return s
}

func (s *Set[T]) Len() uint {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.count
}

func (s *Set[T]) Clear() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = map[T]bool{}
	s.count = 0
}

func (s *Set[T]) Add(item T) {

	if s.Has(item) {
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	s.items[item] = true
	s.count++
}

func (s *Set[T]) Has(item T) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()
	_, exists := s.items[item]
	return exists
}

func (s *Set[T]) Remove(item T) {

	if !s.Has(item) {
		return
	}

	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.items, item)
	s.count--
}

func (s *Set[T]) Items() []T {
	s.lock.RLock()
	defer s.lock.RUnlock()
	items := make([]T, s.Len())
	i := 0
	for item := range s.items {
		items[i] = item
		i++
	}
	return items
}
