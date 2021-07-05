package main

import "fmt"

type account struct {
	balance    int
	Firstname  string
	Secondname string
}

func (a *account) withdrawPointer(amount int) {
	a.balance -= amount
}
func (a account) withdrawValue(amount int) {
	a.balance -= amount
}
func (a account) withdrawReturnValue(amount int) account {
	a.balance -= amount
	return a
}
func main() {
	var test *account = &account{100, "charles", "jeong"}
	test.withdrawPointer(30)
	fmt.Println(test.balance)
	test.withdrawValue(30)
	fmt.Println(test.balance)
	var test2 account = test.withdrawReturnValue(30)
	fmt.Println(test2.balance)
	test.withdrawPointer(10)
	fmt.Printf("test1 : %d\n", test.balance)
	fmt.Printf("test2 : %d\n", test2.balance)

}
