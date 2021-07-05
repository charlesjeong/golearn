package actor

type Actor struct {
}

// func (a *Actor) String(job string) {
// 	fmt.Printf("actor : %s\n", job)
// }

func (a *Actor) String() string {
	return "actor"
}
