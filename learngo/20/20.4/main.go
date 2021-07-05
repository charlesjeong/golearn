package main

import (
	"github.com/charlesjeong/learngo/20/fedex"
	"github.com/charlesjeong/learngo/20/koreaPost"
)

type Sender interface {
	Send(parcel string)
}

func SendBook(name string, sender Sender) {
	sender.Send(name)
}

func main() {
	fedexSender := &fedex.FedexSender{}
	SendBook("어린왕자", fedexSender)
	fedexSender.Send("test")

	koreaSender := &koreaPost.PostSender{}
	SendBook("어린왕자", koreaSender)
	SendBook("소나기", koreaSender)

}

// import (
// 	"github.com/charlesjeong/learngo/20/fedex"
// 	"github.com/charlesjeong/learngo/20/koreaPost"
// )

// type Sender interface {
// 	Send(parcel string)
// }

// func SendBook(name string, sender Sender) {
// 	sender.Send(name)
// }

// func main() {
// 	fedexSender := &fedex.FedexSender{}
// 	SendBook("어린 왕자", fedexSender)
// 	SendBook("그리스인 조르바", fedexSender)

// 	koreaPostSender := &koreaPost.PostSender{}
// 	SendBook("어린 왕자", koreaPostSender)
// 	SendBook("그리스인 조르바", koreaPostSender)

// }
