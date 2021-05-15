package main

import (
	"fmt"
	"src/github.com/charlesjeong/learngo/bank"
)

func main() {
	fmt.Println("helloworld")
	account := bank.MakeAccount("june")
	fmt.Println(account)
}
