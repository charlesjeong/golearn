package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6}
	idx := 2
	//idx2 := 3
	//slice1 = append(slice1[:idx], append([]int{100}, append(slice1[idx:idx2], append([]int{200}, slice1[idx2:]...)...)...)...)
	slice2 := append(slice1, 0)
	copy(slice2[idx+1:], slice1[idx:])
	slice2[idx] = 100
	fmt.Println(slice2)
}
