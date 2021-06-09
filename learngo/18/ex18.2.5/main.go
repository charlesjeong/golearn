package main

import "fmt"

func main() {
	slice1 := make([]int, 3, 3)
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	if cap(slice1) <= len(slice1) {
		slice1 = make([]int, len(slice1), cap(slice1)*2)
		fmt.Println("slice1 expend")
		fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
		fmt.Println("===========")
	}
	slice2 := append(slice1, 4, 5)
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
	slice1 = append(slice1, 100)
	if cap(slice1) <= len(slice1) {
		slice1 = make([]int, len(slice1), cap(slice1)*2)
		fmt.Println("slice1 expend")
		fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
		fmt.Println("===========")
	}
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
	slice1[1] = 100
	fmt.Println("slice1 : ", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2 : ", slice2, len(slice2), cap(slice2))
}
