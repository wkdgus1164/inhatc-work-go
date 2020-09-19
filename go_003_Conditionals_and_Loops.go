package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// 이름은 문자로 시작, 이후에는 문자/숫자 모두 사용 가능
// 변수, 함수, 타입의  이름이 대문자로 시작하면 외부로 노출되어 외부 패키지에서 접근 가능함.
func main() {
	appendText()
}

func stringReplacer() {
	wrongTexts := "Gx Gx Gx!"
	replacer := strings.NewReplacer("x", "o")
	correctTexts := replacer.Replace(wrongTexts)
	fmt.Println(correctTexts)
}

func timeGetter() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}

func math() {
	fmt.Println("반지름 입력 : ")
	reader := bufio.NewReader(os.Stdin)
	in, err := reader.ReadString('\n')
	fmt.Println(in)
	if err != nil {
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)
	radius, err := strconv.ParseFloat(in, 64)
	fmt.Println("원의 넓이 : ", radius*3.141592)
}

func appendText() {
	// error 로 사용할 경우 내장함수 error 에 있어서 Shadowing 이 발생할 수 있음.
	var append string = "Shadow"
	fmt.Println(append)
}
