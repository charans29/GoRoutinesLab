package main

import "fmt"

const (
	IND = "India"
	USA = "USA"

	indiaGreets = "Namastey, "
	usaGreets   = "Hi, "
)

func Hello(name string, region string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(region) + name
}

func greetingPrefix(region string) (prefix string) {
	switch region {
	case IND:
		prefix = indiaGreets
	case USA:
		prefix = usaGreets
	default:
		prefix = "Hello, "
	}
	return
}

func main() {
	fmt.Println(Hello("Charan", "India"))
}
