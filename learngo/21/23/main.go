package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Now()
	fmt.Println(today)

	switch {
	case today.Hour() > 10:
		fmt.Println("atfernoon")
	case today.Hour() > 8:
		fmt.Println("morning")
	default:
		fmt.Println("error")
	}

}
