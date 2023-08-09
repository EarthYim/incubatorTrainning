package main

import (
	"fmt"
	"math"
)

type calType func(float64, float64) float64

var minus = func(a, b float64) float64 {
	return a - b
}

var plus = func(a, b float64) float64 {
	return a + b
}

func compute(fn calType) float64 {
	return fn(3.0, 2.0)
}

func main() {
	r1 := compute(minus)
	r2 := compute(plus)
	fmt.Println(r1, r2)
	r3 := compute(math.Hypot)
	fmt.Println(r3)
}
