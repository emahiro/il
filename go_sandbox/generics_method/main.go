package main

import "fmt"

func main() {
	s := &Stack[int]{}
	s.Push(1)
	s.Push(2)
	s.Push(3)

	mappedStack := s.Map(func(x int) string {
		return "Number: " + fmt.Sprint(x+'0')
	})

	for len(mappedStack.item) > 0 {
		println(mappedStack.Pop())
	}
}

type Stack[T any] struct {
	item []T
}

func (s *Stack[T]) Push(item T) {
	s.item = append(s.item, item)
}

func (s *Stack[T]) Pop() T {
	if len(s.item) == 0 {
		var zero T
		return zero
	}
	item := s.item[len(s.item)-1]
	s.item = s.item[:len(s.item)-1]
	return item
}

func (s *Stack[T]) Map[U any](f func(T) U) *Stack[U] {
	newStack := &Stack[U]{}
	for _, item := range s.item {
		newStack.Push(f(item))
	}
	return newStack
}
