package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, 3, 10)
	slice3 := make([]int, len(slice1))
	cnt1 := copy(slice2, slice1)
	copy(slice3, slice1)

	fmt.Println(cnt1, slice2)
	fmt.Println(slice3)
}
