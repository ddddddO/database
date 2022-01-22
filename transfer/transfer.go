// クライアントからクエリを受け付ける
// TODO: ここからやる
package transfer

import (
	"fmt"
)

type transferer struct {
	sendQueue chan<- string
}

func New(queue chan<- string) *transferer {
	return &transferer{
		sendQueue: queue,
	}
}

func (t *transferer) Run() error {
	t.transfer()
	return nil
}

func (t *transferer) transfer() {
	fmt.Println("Transfer!")

	// netパッケージか何か使ってクライアントからクエリを受け付ける
	dummy := []string{"a", "b", "c"}
	for _, v := range dummy {
		t.sendQueue <- v
	}
}

func (t *transferer) close() {

}
