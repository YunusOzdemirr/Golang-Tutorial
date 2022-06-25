package main

import (
	"errors"
	"fmt"
)

func mai2n() {
	myError := errors.New("Zort error")
	fmt.Println(myError)
}
