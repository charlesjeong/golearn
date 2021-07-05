package main

type Stringer interface{ String() }
type Reader interface{ Read() }

func CheckAndRun(Stringer Stringer) {
	if r, ok := Stringer.(Reader); ok {
		r.Read()
	}
}

// func main() {/fmt.Println("")

// import (
// 	"fmt"

// 	"github.com/charlesjeong/learngo/20/20_test/AwesomeDB"
// )

// type DB interface {
// 	GetData()
// 	WriteData(data string)
// 	Close() string
// }

// func main() {
// 	s := &AwesomeDB.OurDB{"test"}
// 	fmt.Println(s)
// 	fmt.Println(s.Close())
// }

//1번 문제
// type ReadWriter interface {
// 	Read()
// 	Write()
// }
// type File struct{}

// func (f *File) Read()  {}
// func (f *File) Write() {}
// func ReadWrite(rw ReadWriter) {
// 	rw.Read()
// 	fmt.Println("Read")
// 	rw.Write()
// 	fmt.Println("Write")
// }

// func main() {
// 	f := &File{}
// 	ReadWrite(f)
// }
