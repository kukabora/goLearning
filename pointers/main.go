package main

import "fmt"

func main() {
	var A *int
	var B int = 50
	A = &B

	fmt.Println(A)
	fmt.Println(*A)

	*A += 50
	fmt.Println(A)
	fmt.Println(*A)
}
