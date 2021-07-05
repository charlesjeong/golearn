package main

import "fmt"

func main() {
	age := 11

	switch true {
	case age < 20 && age > 10:
		fmt.Println("개좋다")
	case age < 10:
		fmt.Println("좋다")

	default:
		fmt.Println("bug")
	}
}
