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
	basicArray()
}
func basicArray() {
	// myArray := [3]int{}
	// myArray[0] = 1
	// myArray[1] = 2
	// myArray[2] = 3
	// fmt.Println(myArray)

	// colors := [3]string{}
	// colors[0] = "black"
	// colors[1] = "white"
	// colors[2] = "blue"
	// fmt.Println(colors)

	colors := []string{"black", "white", "blue", "red", "green", "yellow"}
	//numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//fmt.Println(len(numbers))
	//fmt.Println(len(colors))
	for i := 0; i < len(colors)-1; i++ {
		//	fmt.Println(colors[i])
	}
	// for i, value := range colors {
	// 	fmt.Println(i, value)
	// }
	for _, value := range colors {
		fmt.Println(value)
	}

}

func forloop() {

	for i := 0; ; i++ {
		fmt.Println(i)
		if i == 5000 {
			break
		}
	}
}

func HumanCreate() {
	var human = NewHuman()
	human.FirstName = "asd"
	human.LastName = "sdsad"
	human.Age = 12

	if human != nil {
		fmt.Println("Vay amk")
	}

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
