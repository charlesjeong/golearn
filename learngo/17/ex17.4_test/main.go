package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func inputNumber() int {
	var number int
	fmt.Printf("input number : (0~10)")
	_, err := fmt.Scanln(&number)
	for err != nil {
		stdin.ReadString('\n')
		fmt.Printf("again input number : (0~10)")
		_, err = fmt.Scanln(&number)
	}
	return number
}
func makeRandomInt() int {
	rand.Seed(time.Now().UnixNano())
	randint := rand.Intn(10)
	return randint
}
func main() {
	var number int
	var rand_int = makeRandomInt()
	fmt.Printf("맞춰야하는 수는 %d 입니다\n", rand_int)
	number = inputNumber()
	for number != rand_int {
		if number > rand_int {
			fmt.Printf("%d 보다 작은 수 입니다.\n", number)
		} else {
			fmt.Printf("%d 보다 큰 수 입니다.\n", number)
		}
		number = inputNumber()
	}
	fmt.Println("정답입니다.")
}
