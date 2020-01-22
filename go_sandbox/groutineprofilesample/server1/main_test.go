package main

import (
	"testing"

	"github.com/pkg/profile"
)

func TestExecUser(t *testing.T) {
	defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	getUser()
}
