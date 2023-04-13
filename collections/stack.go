// this is an implementation of a LIFO data structure
// the goal as well is to keep it thread-safe and generic
package collections

import "sync"

type element[T any] struct {
	Data     T
	Next     *element[T]
	Previous *element[T]
}

type stack[T any] struct {
	head   *element[T]
	tail   *element[T]
	length int
	mutex  sync.Mutex
}

func (s *stack[T]) Push(data T) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	elem := &element[T]{Data: data}

	if s.length == 0 {
		s.head = elem
		s.tail = elem
	} else {
		elem.Previous = s.tail
		s.tail.Next = elem
		s.tail = elem
	}

	s.length++
}

func (s *stack[T]) Peek() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.tail.Data
}

func (s *stack[T]) Pop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.length == 0 {
		return
	}
	if s.length == 1 {
		s.head = nil
		s.tail = nil
	}
	s.tail = s.tail.Previous
	s.tail.Next = nil
	s.length--
}

func New[T any](T datum) *stack[T] {
	return &stack[T]{head: nil, tail: nil, length: 0, mutex: sync.Mutex{}}
}
