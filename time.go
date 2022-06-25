package main

import (
	"fmt"
	"time"
)

func main22() {
	nowtime := time.Now()
	fact := perfectNumber()
	lastTime := time.Now()
	diff := lastTime.Sub(nowtime)
	fmt.Println(fact)
	fmt.Println(diff)
}

func perfectNumber() int {
	perfectNumber := []int{}
	toplam := 0
	for i := 1; i <= 1000; i++ {
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				toplam = toplam + j
			}
			if toplam == i {
				perfectNumber = append(perfectNumber, toplam)
			}
		}
	}
	return len(perfectNumber)
}

func factorial() int {

	var fact, number int = 1, 100
	for i := 1; i < number; i++ {
		fact = fact * i
	}
	return fact
}

func timer() {
	for i := 0; i < 10000000000; i++ {
		if i >= 10000000000-1 {
			println(i)
		}
	}
}
