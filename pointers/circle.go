package main

import (
	"fmt"
	"math"
)

const length = 35

func main() {
	var R *float64
	var raduis float64 = 35 / 2.0 / math.Pi
	R = &raduis

	fmt.Println(math.Pow(*R, 2) * math.Pi)
}
