package main

import "testing"

func TestHello(t *testing.T)  {
	got := Hello("Firas")
	want := "Hello, Firas"

	if got != want {
		t.Errorf("got %q and want %q",got,want)
	}
}