package main

import (
	"fmt"
	"sync"
)

var (
	m  = make(map[string]int64, 0)
	mu = sync.Mutex{}
)

func main() {
	fmt.Printf("test")
}

func race(key string, val int64) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Printf("key = %v, val = %v\n", key, val)
	m[key] = val
	fmt.Printf("m = %v\n", m)
}
