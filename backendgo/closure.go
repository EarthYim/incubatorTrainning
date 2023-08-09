package main

import "fmt"

func adder() (func() int, func() int) {
	sum := 0
	inc := func() int {
		sum = sum + 1
		return sum
	}

	cur := func() int {
		return sum
	}
	return inc, cur
}

func main() {
	inc, cur := adder()
	fmt.Println(cur())

	fmt.Println(inc())
	fmt.Println(inc())

	fmt.Println(cur())
	fmt.Println(cur())

}
