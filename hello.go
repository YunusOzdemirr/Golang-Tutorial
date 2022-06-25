package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	channel := make(chan int)
	go Sum(numbers, channel)
	go Sum(numbers[:len(numbers)/2], channel)
	x, y := <-channel, <-channel
	fmt.Println(x)
	fmt.Println(y)
}
func Sum(numbers []int, c chan int) {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	c <- sum
}

func GoRoutines() {
	runtime.GOMAXPROCS(8)
	go SayHello("Wassup")
	SayHello("Hello")
}
func SayHello(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		go fmt.Println(s)
	}
}

func basicMap() {
	//KeyValuePair
	states := make(map[string]string)
	states["ist"] = "istanbul"
	states["ank"] = "ankara"
	states["ant"] = "antalya"
	fmt.Println(states)
	//How to get value by key
	antalya := states["ant"]
	fmt.Println(antalya)
	//How to delete value by key
	delete(states, "ank")
	fmt.Println(states)

	// for k, v := range states {
	// 	fmt.Println(k, v)
	// }

	keys := make([]string, len(states))
	i := 0
	for k, v := range states {
		keys[i] = k + " " + v
		i++
	}
	fmt.Println(keys)

}

func basicMake() {
	var numbers = make([]int, 5, 5)
	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 4
	fmt.Println(numbers)
	numbers = append(numbers, 232)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))
	fmt.Println(len(numbers))
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

	// for i := 0; i < len(colors)-1; i++ {
	// 	//	fmt.Println(colors[i])
	// }
	// // for i, value := range colors {
	// // 	fmt.Println(i, value)
	// // }
	// for _, value := range colors {
	// 	fmt.Println(value)
	// }

	colors = append(colors, "purple")
	fmt.Println(colors)
	colors = append(colors[0 : len(colors)-1])
	fmt.Println(colors)
	colors = append(colors[1:])
	fmt.Println(colors)

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
func VertextCreate() {
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
}

type Human struct {
	FirstName string
	LastName  string
	Age       int
}

func NewHuman() *Human {
	human := new(Human)
	return human
}

type Vertext struct {
	X int
	Y int
}
