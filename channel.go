package main

import (
	"fmt"
)

func main() {
	k1 := make(chan string)
	k2 := make(chan string)
	go func() {
		k1 <- "video"
	}()

	go func() {
		k2 <- "ses"
	}()
	for i := 0; i < 2; i++ {
		select {
		case mesaj1 := <-k1:
			fmt.Println(mesaj1)
		case mesaj2 := <-k2:
			fmt.Println(mesaj2)
		}
	}
}
