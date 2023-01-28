package main

import "testing"

func TestHelloWorld(t *testing.T) {
	got := Hello()
	want := "Hello, World"

	if got != want {
		t.Errorf("got is not equal to want")
	}
}
