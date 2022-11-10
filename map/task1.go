package main

import (
	"fmt"
	"strconv"
)

type Editions map[string][]string

func main() {
	m := map[string]Editions{
		"User1": Editions{
			"J.Rolling":   []string{"Harry Potter 1", "Harry Potter 2"},
			"Dostoyevkiy": []string{"Prestupleniye i nakazaniye"},
		},
		"User2": Editions{
			"J.Rolling": []string{"Harry Potter 3"},
		},
		"User3": Editions{},
	}

	counter := 0

	sortedData := map[string]int{}

	for k, v := range m {
		sortedData[k] = len(v)
		if len(v) > 0 {
			counter++
		}
	}

	fmt.Println("Amount of users with non-returned books is :" + strconv.Itoa(counter))
	fmt.Println("------------------")
	fmt.Println("Amount of books on each  reader")
	for k, v := range sortedData {
		fmt.Println(k, v)
	}

}
