package main

import "fmt"

type a struct {
	B b `json:"b"`
}

type b struct {
	C c `json:"c"`
}

type c struct{}

func parseJSON(data string) error {

}

func main() {
	fmt.Println("hello")
}
