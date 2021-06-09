package main

import "fmt"

func changeArray(array2 [5]int) {
	array2[4] = 10
}
func changeSlice(slice2 []int) {
	slice2[4] = 10
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}
	slice := []int{1, 2, 3, 4, 5}
	changeSlice(slice)
	changeArray(array)
	fmt.Println("array :", array)
	fmt.Println("slice : ", slice)

}
