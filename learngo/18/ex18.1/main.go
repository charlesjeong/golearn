package main

import "fmt"

func main() {
	//var slice []int = []int{1, 2, 3}
	var slice = make([]int, 3, 5)
	//slice2 := slice
	slice2 := append(slice, 4, 5)
	slice[2] = 100
	//var slice = make([]int, 3, 5)

	//fmt.Println(len(slice))
	fmt.Println(slice, len(slice), cap(slice))
	fmt.Println(slice2, len(slice2), cap(slice2))
}
