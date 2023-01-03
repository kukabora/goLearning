package wordz

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

var Hello = "This is package wordz"
var Prefix = "Random word: "
var Words = []string{"one", "two", "three", "four", "five"}

func init() {
	fmt.Println("Function init in package wordz")
}

func Random(Words []string) string {
	max := len(Words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64(), Words)
}

func get(index int64, Words []string) string {
	return Prefix + Words[index]
}
