package main

import "fmt"

var m = make(map[string]int64, 0)

func main() {
	fmt.Printf("test")
}

func race(key string, val int64) {
	fmt.Printf("key = %v, val = %v\n", key, val)
	m[key] = val
	fmt.Printf("m = %v\n", m)
}
