package main

import "fmt"

func main() {
	name := "Hacker"

	switch name {
	case "Sanya":
		fmt.Println("Hello, Alex")
	case "Igor":
		fmt.Println("Access, Igor")
	default:
		panic("Unknown name!")
	}
}
