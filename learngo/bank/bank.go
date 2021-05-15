package bank

import "fmt"

type account struct {
	owner   string
	balance int
}

func MakeAccount(owner string) *account {
	account := account{owner: owner}
	return &account
}

func Test() {
	fmt.Println("test")
}
