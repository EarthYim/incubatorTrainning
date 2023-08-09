package main

import (
	"fmt"
)

const (
	sunday int = iota
	monday
	tuesday
	wednesday = iota
	thursday
	friday
	saturday
)

func printDay(s string) {
	fmt.Println(s)
}

func main() {
	// fmt.Println(sunday)
	// fmt.Println(wednesday)
	// fmt.Println(fz.Fizzbuss(33))

	i := 42
	j := 33
	fmt.Println("i: ", i)

	var p *int
	fmt.Println("p: ", p)
	fmt.Printf("i: %p\n", &i)

	p = &i
	fmt.Println("p: ", p)

	i = 55
	fmt.Println()
	fmt.Println("p:", *p)

	*p = 33
	fmt.Println("new i: ", i)

}
