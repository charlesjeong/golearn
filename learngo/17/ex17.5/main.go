package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() int {
	var n int
	a, err := fmt.Scanln(&n)
	for err != nil {
		stdin.ReadString('\n')
		fmt.Printf("%d %d\n", a, err)
		fmt.Println("숫자만 입력하시오")
		a, err = fmt.Scanln(&n)
	}
	fmt.Println("입력한 숫자는", n, "입니다")
	return n
}
func MakeRandNum() int {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(10)
	return n
}
func main() {
	var n int
	var randn int
	randn = MakeRandNum()
	fmt.Printf("숫자를 입력하세요(0~9) : ")
	fmt.Println("정답 :", randn)
	n = InputIntValue()
	for n != randn {
		if n > randn {
			fmt.Println(n, "보다 작습니다.")
		} else {
			fmt.Println(n, "보다 큽니다.")
		}
		fmt.Printf("다시 입력 : ")
		n = InputIntValue()
	}
	fmt.Println("맞췄습니다.")
}
