package main

import (
	"fmt"
)


func main() {
	str:="hello월드"
	runes := []rune (str)

	for i:=0; i<len(runes);i++ {
		fmt.Printf("type:%T string:%c\n",runes[i],runes[i])
	}
	//회사테스트
}
