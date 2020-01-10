package main

import (
	"testing"

	"github.com/pkg/profile"
)

func TestExecUser(t *testing.T) {
	defer profile.Start(profile.GoroutineProfile, profile.ProfilePath(".")).Stop()

	cnt := 3
	for i := 0; i < cnt; i++ {
		go func() {
			if _, err := getUser(); err != nil {
				t.Fatal(err)
			}
		}()
	}
}
