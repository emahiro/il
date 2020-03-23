package main

import "testing"

func TestEchoA(t *testing.T) {
	arg := "hello"
	got := EchoA(arg)
	if got != arg {
		t.Fatalf("got is %s", got)
	}
}

func TestX_EchoX(t *testing.T) {
	arg := "hello"
	x := X{}
	got := x.EchoX(arg)
	if got != arg {
		t.Fatalf("got is %s", got)
	}
}
