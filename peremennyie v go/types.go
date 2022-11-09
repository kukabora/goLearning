package main

import (
	"fmt"
	"math"
)

type AmericanVelocity float64
type EuropeanVelocity float64

func main() {
	var x AmericanVelocity = AmericanVelocity(math.Round(120.4*3.6*100) / 100)
	var y EuropeanVelocity = EuropeanVelocity(math.Round((130*3600.0)/1609*100) / 100)
	fmt.Println(x)
	fmt.Println(y)
}
