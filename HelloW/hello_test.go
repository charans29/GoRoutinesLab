package main

import "testing"

func TestHelloW(t *testing.T){
	got := Hello()
	want := "Hello, Universe"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}