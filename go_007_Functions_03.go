package main

import (
	"fmt"
	"reflect"
)

func main() {
	id := 55
	grade := 3.91
	pid := &id
	pgrade := &grade
	fmt.Println(grade, &grade, reflect.TypeOf(grade))
	fmt.Println(*pgrade, &grade, reflect.TypeOf(pgrade))
	fmt.Println(id, &id)
	fmt.Println(*pid, pid, &pid)
	//number := 3
	//cube(number)
}

//func cube(n int) {
//	n = n * n * n
//	fmt.Println(n)
//}
