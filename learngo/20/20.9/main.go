package main

import (
	"fmt"

	"github.com/charlesjeong/learngo/20/actor"
	"github.com/charlesjeong/learngo/20/student"
)

// type Stringer interface {
// 	String(job string)
// }

// func SendJob(name string, stringer Stringer) {
// 	stringer.String(name)
// }

type Stringer interface {
	String() string
}

func ConvertType(stringer Stringer) {
	s, ok := stringer.(*student.Student)
	if ok {
		fmt.Printf("%T : %v", stringer, s.String())
	} else {
		fmt.Printf("%T to *student.Student Convert false\n", stringer)
	}
	// switch t := stringer.(type) {
	// case *student.Student:
	// 	s, ok := stringer.(*student.Student)
	// 	fmt.Println(s.String(), ok)
	// case *actor.Actor:
	// 	a, ok := stringer.(*actor.Actor)
	// 	fmt.Println(a.String(), ok)
	// default:
	// 	fmt.Printf("Not Supported Type : %T", t)
	// }
}

// original
// student := stringer.(*student.Student)
// fmt.Println(student)
// }
func main() {
	// ConvertType에서 *Actor가 *Student 타입으로 변환할려고 하니 안된다.
	actor := &actor.Actor{}
	ConvertType(actor)
	student := &student.Student{}
	ConvertType(student)
	// actor := &actor.Actor{}
	// SendJob("test", actor)
	// student := &student.Student{}
	// SendJob("test2", student)
}
