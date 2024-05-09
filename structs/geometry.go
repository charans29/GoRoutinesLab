package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}