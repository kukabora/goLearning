package taskModule

import (
	"fmt"

	wordz "sample.com/wordz"
)

var digits = []string{"One", "Two", "Three", "Four"}

func Digit() {
	fmt.Println(wordz.Random(digits))
}
