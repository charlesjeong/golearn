package main

import (
	"fmt"
	"time"
)

//1번문제
// type Cart struct {
// 	productList string
// }

// func (c *Cart) AddProduct(product string) {
// 	if c.productList != "" {
// 		c.productList += ", "
// 	}
// 	c.productList += product
// }
// func (c *Cart) Clear() {
// 	c.productList = ""
// }
// func (c Cart) GetProductList() string {
// 	return c.productList
// }

// type ParkingLot struct {
// 	LotSize int
// }
//2번문제
// func ParkCar(lot *ParkingLot, carSize int) {
// 	lot.LotSize -= carSize
// }
// func (lot *ParkingLot) ParkCar(carSize int) {
// 	lot.LotSize -= carSize
// }
// type myString string

// func (m myString) ToLower() myString {
// 	str := strings.ToLower(string(m))
// 	return myString(str)
// }
// func (m myString) ToUpper() myString {
// 	str := strings.ToUpper(string(m))
// 	return myString(str)
// }
type Courier struct {
	Name string
}
type Product struct {
	Name  string
	Price int
	ID    int
}
type Parcel struct {
	Pdt           *Product
	ShippedTime   time.Time
	DeleveredTime time.Time
}

func (c *Courier) SendProduct(pdt *Product) *Parcel {
	test := Parcel{}
	test.Pdt = pdt
	test.ShippedTime = time.Now()
	return &test
}

func (p *Parcel) Delivered() *Product {
	p.DeleveredTime = time.Now()
	return p.Pdt
}
func main() {
	cname := Courier{"CJ"}
	pdt := Product{
		Pdt := "usb"
	}
	//3번 문제
	// msg := myString("hello Go World")
	// fmt.Println(msg)
	// msg2 := msg.ToLower()
	// fmt.Println(msg2)
	// msg3 := msg.ToUpper()
	// fmt.Println(msg3)

	//2번문제
	// lot := &ParkingLot{100}
	// ParkCar(lot, 10)
	// lot.ParkCar(10)
	// fmt.Println(lot)
	//1번 문제
	// c := &Cart{}
	// c.AddProduct("apple")
	// c.AddProduct("kimchi")
	// fmt.Println(c.GetProductList())
	// c.Clear()
	// c.AddProduct("watermelon")
	// fmt.Println(c.GetProductList())
}
