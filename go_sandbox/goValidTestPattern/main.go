package main

import "fmt"

func main() {
	fmt.Printf("hello")
}

// EchoA ...
func EchoA(a string) string {
	return a
}

// X ...
type X struct{}

// EchoX ...
func (x X) EchoX(xx string) string {
	return xx
}
