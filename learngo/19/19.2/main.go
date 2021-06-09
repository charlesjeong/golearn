package main

import (
	"fmt"
	"sort"
)

type Atheletic struct {
	Name              string
	Age               int
	Score             int
	Pass_success_rate float64
}
type Atheletics []Atheletic

func (a Atheletics) Len() int           { return len(a) }
func (a Atheletics) Less(i, j int) bool { return a[i].Pass_success_rate < a[j].Pass_success_rate }
func (a Atheletics) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	A := []Atheletic{
		{"나통키", 13, 45, 78.4},
		{"오맹태", 16, 24, 67.4},
		{"오동도", 18, 54, 50.8},
		{"황금산", 16, 36, 89.7},
	}
	sort.Sort(sort.Reverse(Atheletics(A)))
	fmt.Println(A)
	//P385
	//2번 문제
	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slice = slice[:len(slice)-2]
	// fmt.Println(slice)
	//3번 문제
	// slice := []int{1, 2, 3, 4, 5, 6}
	// t, slice := slice[len(slice)-1], slice[:len(slice)-1]
	// fmt.Println(t, slice)
	//4번 문제
}
