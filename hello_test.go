package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Harry")
	want := "Hello, Harry"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}