package main

import "fmt"

func main() {
	// i := 0
	// for {
	// 	fmt.Println(i)
	// 	i++
	// }

	// loop1:
	// 	for {
	// 		for i := 0; i < 10; i++ {
	// 			fmt.Println(i)
	// 			if i == 5 {
	// 				break loop1
	// 			}
	// 		}
	// 	}

	// flag := true
	// for flag {
	// 	fmt.Println("Hello")
	// 	flag = false
	// }

	// Перебор массивов и мапов

	// arr := []int{1, 2, 3}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr := []int{1, 2, 3}
	// for k, v := range arr {
	// 	fmt.Println(k, v)
	// }

	var custom_map map[string]string = map[string]string{
		"Key1": "value1",
		"Key2": "value2",
	}

	for k, v := range custom_map {
		fmt.Println(k, v)
	}
}
