package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	// 1. make 및 순회해서 값 복사하는 방법
	// slice2 := make([]int, len(slice1))
	// for i, v := range slice1 {
	// 	slice2[i] = v
	// }

	// 2. append를 써서 복사하는데 값이 많으면 노가다이다..
	// slice2 := append([]int{}, slice1[0], slice1[1], slice1[2], slice1[3], slice1[4])

	// 3. make copy를 써서 복사 2줄이면 된다.
	slice2 := make([]int, len(slice1))
	fmt.Println(slice2)
	copy(slice2, slice1)
	slice1[0] = 100
	fmt.Println(slice2)

}
