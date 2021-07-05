package student

type Student struct {
}

// func (a *Student) String(job string) {
// 	fmt.Printf("student : %s\n", job)
// }
func (s *Student) String() string {
	return "student"
}
