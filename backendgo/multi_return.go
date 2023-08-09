package main

import "fmt"

func swap(a, b int) (x, y int) {
	x = b
	y = a
	return
}

func main() {
	x := 10
	y := 5
	x, y = swap(x, y)
	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
}
