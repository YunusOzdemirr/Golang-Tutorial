package main

import (
	"fmt"
	"strconv"
)

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func NewHuman() *Human {
	human := new(Human)
	return human
}

var sayi = 5

func add(x int, y int, z string) int {
	return x + y
}

func multiple(x int, y int) int {
	return x * y
}
func NewVertex() *Vertext {
	h := new(Vertext)
	return h
}

func main() {
	// fmt.Println(add(5, 10))
	// fmt.Println(multiple(5, 10))
	// var message string = "Hi"
	// PrintConsole(&message)
	// fmt.Println(message)

	// v := Vertext{Y: 3}
	// a := &v
	// fmt.Println(a)
	// v.X = 010e9
	// fmt.Println(v)
	var human = NewHuman()
	human.FirstName = "asd"
	human.LastName = "sdsad"
	human.Age = 12

	var str string = human.FirstName + " " + human.LastName + " " + strconv.Itoa(human.Age)
	fmt.Println(str)
}

func PrintConsole(mess *string) {
	// fmt.Println(mess)
	var str = string(*mess)
	fmt.Println(str)
}

type Vertext struct {
	X int
	Y int
}
