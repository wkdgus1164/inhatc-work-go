package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//randomNumber()
	fmt.Println(fact(10))
}

func randomNumber() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println(target)
}

func fact(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result = result * 1
	}
	return result
}
