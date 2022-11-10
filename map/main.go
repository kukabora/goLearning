package main

import "fmt"

func main() {
	m := map[string]int{"first": 1, "second": 2}

	// third, ok := m["third"]

	// if !ok {
	// 	fmt.Println("Key does not exist") // второе значение позволяет проверить есть ли такой ключ
	// 	return
	// }
	// fmt.Println(third)

	// delete(m, "second")
	// fmt.Println(m)

	m["third"] = 3

	// fmt.Println(m)

	for k, v := range m {
		fmt.Println(k, v)
	}
}
