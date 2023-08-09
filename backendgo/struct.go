package main

import "fmt"

type Runner interface {
	Run()
}

type Cat struct {
	name string
}

type User struct {
	name   string
	age    int
	mobile []string
}

func (c Cat) Run() { // interface
	fmt.Println("running", c.name)
}

func (u User) Run() { // interface
	fmt.Println("walking", u.name, u.age, u.mobile)
}

func (u User) walk() {
	u.mobile[2] = "Nokia"
	fmt.Println("walking", u.name, u.age, u.mobile)
}

func Runable(r Runner) {
	fmt.Println("Runalbe:: ")
	r.Run()
}

func main() {
	u := User{name: "Ea", age: 24, mobile: []string{"iphone", "samsung", "huawei"}}
	u.walk()
	fmt.Println("main:", u)
	fmt.Println(u.mobile)
	c := Cat{name: "Cake"}

}
