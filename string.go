package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var source = rand.NewSource(time.Now().UnixNano())

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func main() {
	password := GeneratePassword(10)
	fmt.Println(password)
}
func GeneratePassword(lenght int) string {
	x := make([]byte, lenght)
	for i := range x {
		x[i] = charset[source.Int63()%int64(len(charset))]
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
