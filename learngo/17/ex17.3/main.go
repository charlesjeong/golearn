package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	test := rand.Intn(10)
	fmt.Println(test)
}
