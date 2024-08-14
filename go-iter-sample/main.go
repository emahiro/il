package main

import (
	"fmt"

	"github.com/emahiro/il/go-iter-sample/list"
)

func main() {
	var l list.List[int]
	l.Push(1)
	l.Push(2)
	l.Push(3)

	for list := range l.All() {
		fmt.Println(list.Value())
	}
}
