package main

import (
	"fmt"
	"strconv"
)

func main() {
	strNumb := 35
	numbStr := "104"
	fmt.Printf("Before convertation(strNumb): %T\n", strNumb)
	fmt.Printf("Before convertation(numbStr): %T\n", numbStr)

	convertedNum := strconv.Itoa(strNumb)
	convertedStr, _ := strconv.Atoi(numbStr)

	fmt.Printf("After convertation(strNumb): %T\n", convertedNum)
	fmt.Printf("After convertation(numbStr): %T\n", convertedStr)
	fmt.Println("---------")
	fmt.Println("Values:")
	fmt.Println(convertedNum)
	fmt.Println(convertedStr)
}
