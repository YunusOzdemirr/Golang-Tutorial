package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Option struct {
	length    int
	upperCase bool
	lowerCase bool
	numbers   bool
	special   bool
}

var source = rand.NewSource(time.Now().UnixNano())

// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
const charsetLowerCase = "abcdefghijklmnopqrstuvwxyz"
const charsetUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumber = "01234567890"
const chartsetSpecial = "!*-)(/&%+^"

func main() {
	password, err := GeneratePasswordStruct(Option{length: 10, upperCase: false, lowerCase: false, numbers: false, special: false})
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(password)
	}
}

func GeneratePasswordStruct(opt Option) (string, error) {
	if opt.length <= 0 {
		return "Lenght is must be greater than zero", errors.New("lenght must be greater than zero")
	}
	charSet := ""
	if opt.lowerCase {
		charSet += charsetLowerCase
	}
	if opt.upperCase {
		charSet += charsetUpperCase
	}
	if opt.numbers {
		charSet += charsetNumber
	}
	if opt.special {
		charSet += chartsetSpecial
	}

	if charSet == "" {
		return "Non Generated", errors.New("Karakter seçimini true yapiniz.")
	}
	x := make([]byte, opt.length)

	for i := range x {
		x[i] = charSet[source.Int63()%int64(len(charSet))]
	}
	return string(x), nil
}

func GeneratePassword(lenght int) string {
	x := make([]byte, lenght)
	for i := range x {
		x[i] = charsetLowerCase[source.Int63()%int64(len(charsetLowerCase))]
	}
	return string(x)
}

func StringBuilderCompare() string {

	var x bytes.Buffer

	for i := 0; i < 10; i++ {
		x.WriteString(RandomString())
	}
	fmt.Println(x.String())
	builder := strings.Builder{}
	for j := 0; j < 10; j++ {
		builder.WriteString("Data")
	}
	result := builder.String()
	fmt.Println(result)
	return result
}

func RandomString() string {
	return "xyz-4+"
}
