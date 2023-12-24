package main

import "testing"

func TestHello(t *testing.T){
	t.Run("saying hello to people",func(t *testing.T) {
		got := Hello("Waheb","")
		want := "Hello, Waheb"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("say 'Hello, World' When an empty sting is supplied",func(t *testing.T) {
		got := Hello("","")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish",func(t *testing.T) {
		got :=Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French",func(t *testing.T) {
		got := Hello("Waheb","French")
		want := "Bonjour, Waheb"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}