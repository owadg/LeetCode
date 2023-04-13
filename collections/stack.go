// this is an implementation of a LIFO data structure
// the goal as well is to keep it thread-safe and generic
package collections

import "sync"

type element[T any] struct {
	Data     T
	Next     *element[T]
	Previous *element[T]
}

type Stack[T any] struct {
	head   *element[T]
	tail   *element[T]
	length int
	mutex  sync.Mutex
}

func (s *Stack[T]) Push(data T) *Stack[T] {
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
	return s
}

func (s *Stack[T]) Peek() T {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	return s.tail.Data
}

func (s *Stack[T]) Pop() *Stack[T] {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.length == 0 {
		return s
	}
	if s.length == 1 {
		s.head = nil
		s.tail = nil
		s.length--
		return s
	}

	s.tail = s.tail.Previous
	s.tail.Next = nil
	s.length--
	return s
}

func NewStack[T any](datum T) *Stack[T] {
	return (&Stack[T]{head: nil, tail: nil, length: 0, mutex: sync.Mutex{}}).Push(datum)
}
