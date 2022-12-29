package main

import "fmt"

func print() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("HELLO WORLD")
}

func main() {
	print()
}
