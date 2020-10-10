package main

import "fmt"

func main() {
	var w, h, area float64
	w = 4.1
	h = 3.0
	area = w * h
	fmt.Printf("%.2f\n", area/10.0)
	w = 5.2
	h = 3.5
	area = w * h
	fmt.Printf("%.2f\n", area/10.0)
	fmt.Printf("%v %#v %T %%\n", 1.2, 1.2, 1.2)
}
