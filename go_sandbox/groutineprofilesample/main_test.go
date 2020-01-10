package main

import (
	"testing"
)

func TestExecUser(t *testing.T) {
	cnt := 3
	for i := 0; i < cnt; i++ {
		go func() {
			if _, err := getUser(); err != nil {
				t.Fatal(err)
			}
		}()
	}
}
