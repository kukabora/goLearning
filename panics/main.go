package main

import "fmt"

func someFunc() {
	defer fmt.Println("somefunc: defer")

	panic("CRITICAL ERROR")
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			switch err {
			case "CRITICAL ERROR":
				fmt.Println("SO good")
			default:
				panic("FATAL ERROR")
			}
		}
	}()
	someFunc()
}
