package main

import (
	"bufio"
	"fmt"
<<<<<<< HEAD
	"math/rand"
	"os"
	"time"
)

func input_number() int {
	var input int
	var stdin = bufio.NewReader(os.Stdin)
	for _, err := fmt.Scanln(&input); err != nil; {
		stdin.ReadString('\n')
		fmt.Println("숫자만 입력하세요")
		_, err = fmt.Scanln(&input)
	}
	return input
}
func make_rand_num() int {
	rand.Seed(time.Now().UnixNano())
	find_number := rand.Intn(100)
	return find_number
}
func main() {

	var input int
	find_number := make_rand_num()

	fmt.Print("(0~99)숫자를 입력하세요 : ")
	input = input_number()
	println(find_number)
	for input != find_number {
		if input > find_number {
			fmt.Printf("%d 보다 작은 수 입니다. 다시 숫자 입력 : ", input)
			input = input_number()
		} else {
			fmt.Printf("%d 보다 큰 수 입니다. 다시 숫자 입력 : ", input)
			input = input_number()
		}
	}
	fmt.Println("맞췄습니다.")

=======

	"github.com/guptarohit/asciigraph"
)


func main() {
	data := []float64{8.17, 8.63, 8.65, 8.92, 8.95, 9.22}
	graph := asciigraph.Plot(data)
	fmt.Println(graph)
>>>>>>> e57cf506da58f611ea75d22e44b20889fe7ed63a
}
