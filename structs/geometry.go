package main

func Perimeter(width float64, height float64) float64 {
	 pm := 2 * (width + height)
     return pm
}

func Area(width float64, height float64) float64 {
	 ar := width * height
	 return ar
}

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimtr(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Ar(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}