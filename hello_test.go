package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("You!")
	want := "Hello You!"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
