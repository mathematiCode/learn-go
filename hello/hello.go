package main

import (
	"fmt"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var greeting = ""
	switch language {
	case "Spanish":
		greeting = "Hola, "
	case "French":
		greeting = "Bonjour, "
	default:
		greeting = "Hello, "
	}

	return greeting + name
}

func main() {
	fmt.Println(Hello("Chris", "English"))
}
