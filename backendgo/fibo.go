package main

import "fmt"

func fibo() func() int {
	sum := 1
	lastVal := 0
	return func() int {
		holder := sum + lastVal
		lastVal = sum
		sum = holder
		return sum
	}
}

func main() {
	object := fibo()
	for i := 0; i < 100; i++ {
		fmt.Println(object())
	}
}
