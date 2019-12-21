package main

import (
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	itr := []int64{1, 2, 3}
	wg := sync.WaitGroup{}
	wg.Add(len(itr))
	for _, i := range itr {
		i := i
		go func(k string, v int64) {
			race(k, v)
			wg.Done()
		}("test", i)
	}
	wg.Wait()
}
