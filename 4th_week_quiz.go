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
	// declare two input numbers (initializes with 0 automatically)
	var firstNumber, secondNumber int

	// get first number
	fmt.Print("첫 번째 정수 입력 : ")

	// save first number
	firstNumber = getNumber()

	// get second number
	fmt.Print("두 번째 정수 입력 : ")

	// save second number
	secondNumber = getNumber()

	// start for loop with 1 ~ 50
	for i := 1; i <= 50; i++ {

		// if both judgements are valid
		if divideNumber(firstNumber, i) == true && divideNumber(secondNumber, i) == true {

			// print that it makes with both of the judges!
			fmt.Println(firstNumber, "와 ", secondNumber, "의 공배수")

			// if first judgement is valid
		} else if divideNumber(firstNumber, i) == true {

			// print that it makes with first number!
			fmt.Println(firstNumber, "의 배수")

			// if second judgement is valid
		} else if divideNumber(secondNumber, i) == true {

			// print that it makes with second number!
			fmt.Println(secondNumber, "의 배수")

			// if both judgements neither valid
		} else {

			// just print it!
			fmt.Println(i)
		}
	}
}

func getNumber() int {
	reader := bufio.NewReader(os.Stdin)

	// Read string data
	in, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// Trim string space
	in = strings.TrimSpace(in)

	// Parse to int
	result, err := strconv.ParseInt(in, 10, 32)
	if err != nil {
		log.Fatal(err)
	}

	// return result
	return int(result)
}

func divideNumber(judgeNumber int, inputNumber int) bool {

	// if division makes
	if inputNumber%judgeNumber == 0 {
		return true

		// if division doesn't makes
	} else {
		return false
	}
}
