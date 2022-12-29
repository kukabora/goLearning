package main

import "fmt"

func getMax(a []int) int {
	max := a[0]
	for _, v := range a {
		if max < v {
			max = v
		}
	}
	return max
}

func main() {
	arr := []int{-241, -6, -1234, -564, -873, -48515, -4815}
	fmt.Println(getMax(arr))
}
