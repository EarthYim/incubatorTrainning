package main

import "fmt"

func sum(xs ...int) int {
	var sum int
	// for i := 0; i < len(xs); i++ {
	// 	sum += xs[i]
	// }
	for _, v := range xs {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 1, 2, 2, 1, 1))
}
