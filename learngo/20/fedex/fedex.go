package fedex

import "fmt"

type FedexSender struct {
}

func (f *FedexSender) Send(parcel string) {
	fmt.Printf("Fedex send %v parcel\n", parcel)
}

// type FedexSender struct {
// }

// func (f *FedexSender) Send(parcel string) {
// 	fmt.Printf("Fedex send %v parcel\n", parcel)
// }
