package main

import (
	"fmt"
	"os"
)

func main() {
	uName := os.Getenv("USER")
	fmt.Println(uName)
}

func GetAllEnvNames() []string {
	envs := []string{}
	i := 0
	for _, env := range os.Environ() {
		envs = append(envs, env)
		fmt.Println(env)
		i++
	}
	return envs
}
