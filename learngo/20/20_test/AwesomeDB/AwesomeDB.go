package AwesomeDB

import "fmt"

type OurDB struct {
	Name string
}

func (db *OurDB) GetData() {
	fmt.Println("GetData")
}
func (db *OurDB) WriteData(data string) {
}
func (db *OurDB) Close() string {
	return "Close call"
}
