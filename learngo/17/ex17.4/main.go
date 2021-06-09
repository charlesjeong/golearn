package main

import (
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)
var sc = bufio.NewScanner(os.Stdin)

func InputIntValue() (int, error) {
	fmt.Println("숫자를 입력하시오")
	var n int
	_, err := fmt.Scanln(&n)
	if err != nil {
		stdin.ReadString('\n')
	}
	return n, err
}
func main() {
	for {
		n, err := InputIntValue()
		if err != nil {
			fmt.Println("숫자만 입력하시오")
		} else {
			fmt.Println("입력한 숫자는", n, "입니다.")
		}
	}
}
