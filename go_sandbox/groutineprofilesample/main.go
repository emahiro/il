package main

import (
	"fmt"

	"github.com/pkg/profile"
)

func main() {
	defer profile.Start(profile.GoroutineProfile, profile.ProfilePath(".")).Stop()
	fmt.Println("hello")
}
