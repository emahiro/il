package list

import (
	"fmt"
	"iter"
)

type List[T any] struct {
	value T
	next  *List[T]
}

func (l *List[T]) Push(v T) {
	list := l
	fmt.Printf("v: %+v next:%+v \n", v, list.next)
	for list.next != nil {
		list = list.next
	}
	list.next = &List[T]{value: v}
}

func (l *List[T]) Value() T {
	return l.value
}

func (l *List[T]) Next() *List[T] {
	return l.next
}

func (l *List[T]) All() iter.Seq[List[T]] {
	return func(yield func(List[T]) bool) {
		for l.next != nil {
			l = l.next
			if !yield(*l) {
				return
			}
		}
	}
}
