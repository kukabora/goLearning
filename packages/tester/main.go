package main

import (
	"fmt"

	"sample.com/wordz"
)

func main() {
	fmt.Println("Hello world!")
	for i := 0; i < 5; i++ {
		fmt.Println(wordz.Hello)
		fmt.Println(wordz.Random())

	}
}
