package main

import "fmt"

func sum(test ...int) int {
	sum := 0
	for _, v := range test {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(0))
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2, 3))
}
