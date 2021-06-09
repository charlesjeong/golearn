package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 4 //5을 제거할거임
	// for i := idx + 1; i < len(slice1); i++ {
	// 	slice1[i-1] = slice1[i]
	// }
	// slice1 = slice1[:len(slice1)-1]
	slice1 = append(slice1[:idx], slice1[idx+1:]...)
	fmt.Println(slice1)
}
