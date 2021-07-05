package main

import "fmt"

type Stringer interface {
	String() string
}
type Student struct {
	Age int
}

func (s Student) String() string {
	return fmt.Sprintf("student age is %d", s.Age)
}
func PrintAge(stringer Stringer) {
	s := stringer.(*Student)
	fmt.Printf("PrintAge : %d\n", s.Age)
}
func main() {
	s := &Student{15}
	PrintAge(s)
}
