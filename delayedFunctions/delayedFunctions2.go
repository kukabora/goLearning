package main

import "fmt"

func print(i int) {
	fmt.Println(i)
}

func main() {
	i := 0

	defer func() {
		print(i)
	}()

	i++
	fmt.Println(i)
}
