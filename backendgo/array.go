package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(s[1:4])
	fmt.Println(s[0:4])
	fmt.Println(s[1:])
	s = append(s, 7)
	fmt.Println(s[1:])
}
