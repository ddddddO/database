// クライアントからクエリを受け付ける
// TODO: ここからやる
package transfer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/ddddddO/database/model"
)

type transferer struct {
	sendQueue chan<- *model.TaskAndConn
}

func New(queue chan<- *model.TaskAndConn) *transferer {
	return &transferer{
		sendQueue: queue,
	}
}

func (t *transferer) Run(_ context.Context) error {
	t.transfer()
	return nil
}

func (t *transferer) transfer() {
	fmt.Println("Transfer!")

	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Printf("cannot listen: %v\n", err)
		return
	}

	// NOTE: 以下のコードの大部分はgodashから持ってきた
	// 接続を待ち受け続ける
	for {
		// 1接続分
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("cannot accept: %v\n", err)
			continue
		}
		log.Println("connected")

		// 複数の接続を扱うためgoroutine
		go func() {
			log.Println("received task")

			receivedTask := model.Task{}
			if err := json.NewDecoder(conn).Decode(&receivedTask); err != nil {
				log.Println(err)
				return
			}
			receivedTask.State = model.StateRecivedQuery

			task := &model.TaskAndConn{
				Task: receivedTask,
				Conn: conn,
			}

			t.sendQueue <- task
		}()
	}

}

func (t *transferer) Close() {
	close(t.sendQueue)
}
