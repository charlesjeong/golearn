package main

import (
	"fmt"
	"unsafe"
)

type Padding struct {
	A string
	B int
	C float64
}

func main() {
	a := Padding{}
	fmt.Println(unsafe.Sizeof(a))
	//회사테스트
}
