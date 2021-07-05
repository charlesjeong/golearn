package main

type Attaker interface {
	Attack()
}

func main() {
	var att Attaker
	att.Attack()
}
