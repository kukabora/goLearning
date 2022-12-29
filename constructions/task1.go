package main

import "fmt"

func contains(a []string, x string) bool {
	flag := false
	for _, v := range a {
		if v == x {
			flag = true
		}
	}
	return flag
}

func main() {
	arr := []string{"Hello", "World", "Yessir", "Test", "Go", "Learning"}
	fmt.Println(contains(arr, "Go"))
}
