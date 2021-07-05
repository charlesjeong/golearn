package main

import "fmt"

type Stringer interface {
	String() string
}
type Student struct {
	Name string
	Age  int
}

func (s Student) String() string {
	return fmt.Sprintf("Name : %s Age : %d", s.Name, s.Age)
}
func PrintAge(stringer Stringer) {
	s := stringer.(Student)
	fmt.Printf("PrintAge : %d", s.Age)
}
func main() {
	s := Student{"charles", 15}
	PrintAge(s)

}
