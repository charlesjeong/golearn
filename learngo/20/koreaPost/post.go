package koreaPost

import "fmt"

type PostSender struct {
}

func (k *PostSender) Send(parcel string) {
	fmt.Printf("koreaPost에서 %s 를 보냅니다. parcel\n", parcel)
}

// type PostSender struct{}

// func (k *PostSender) Send(parcel string) {
// 	fmt.Printf("우체국에서 택배 %v를 보냅니다\n", parcel)
// }
