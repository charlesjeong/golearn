package main

import (
	"fmt"
)

type account struct {
	balance int
}

func (a *account) withdrawMethod(amount int) {
	a.balance -= amount
}
func withdrawFunc(a *account, amount int) {
	a.balance -= amount
}
func main() {
	a := &account{100}
	a.withdrawMethod(30)
	withdrawFunc(a, 30)
	fmt.Println(a.balance)
}
