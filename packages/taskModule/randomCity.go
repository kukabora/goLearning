package taskModule

import (
	"fmt"

	wordz "sample.com/wordz"
)

var cities = []string{"Almaty", "Moskow", "Rayong", "Bangkok"}

func City() {
	fmt.Println(wordz.Random(cities))
}
