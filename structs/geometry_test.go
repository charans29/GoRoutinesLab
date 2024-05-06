package main

import "testing"

func TestPerimeter(t *testing.T){
	got := Perimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPArea(t *testing.T){
	got := Area(10.0, 10.0)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestPerimeterStrt(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimtr(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestAreaStrt(t *testing.T) {
	rectangle := Rectangle{12.0, 6.0}
	got := Ar(rectangle)
	want := 72.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
