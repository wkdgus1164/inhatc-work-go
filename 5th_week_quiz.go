package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	for {
		var menu int

		fmt.Print("1) 절대값 2) 팩토리얼 3) 피보나치 4) 거듭제곱 5) 종료 : ")

		menu = getMenu()
		judgeMenu(menu)
	}
}

func judgeMenu(m int) {

	if m == 1 {
		fmt.Print("정수 입력(절대값 계산) : ")
		number := getInput()
		fmt.Println(absoluteNumber(number))
	} else if m == 2 {
		fmt.Print("정수 입력 (팩토리얼 계산) : ")
		number := getInput()
		fmt.Println(factorialNumber(number))
	} else if m == 3 {
		fmt.Print("정수 입력 (피보나치 출력) : ")
		number := getInput()
		fmt.Println(factorialNumber(number))
	} else if m == 4 {
		fmt.Print("Base 입력 : ")
		number := getInput()
		fmt.Print("Exponent 입력 : ")
		number2 := getInput()
		fmt.Println(multipleNumbers(number, number2))
	} else if m == 5 {
		os.Exit(3)
	}
}

func getMenu() int {
	reader := bufio.NewReader(os.Stdin)

	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)

	result, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	return int(result)
}

func getInput() int {
	reader := bufio.NewReader(os.Stdin)

	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	in = strings.TrimSpace(in)

	result, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	return int(result)
}

func absoluteNumber(n int) int {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

func factorialNumber(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorialNumber(n-1)
}

func multipleNumbers(n int, m int) int {
	var i int
	var o int = 1
	for i = 1; i <= m; i++ {
		o = o * n
	}
	return o
}
