package query_processor

import (
	"context"
	"fmt"

	"github.com/ddddddO/database/model"
	"github.com/ddddddO/database/query_processor/optimizer"
	"github.com/ddddddO/database/query_processor/parser"
)

type queryProcessor struct {
	recieveQueue <-chan *model.TaskAndConn
	sendQueue    chan<- string
}

func New(rQueue <-chan *model.TaskAndConn, sQueue chan<- string) *queryProcessor {
	return &queryProcessor{
		recieveQueue: rQueue,
		sendQueue:    sQueue,
	}
}

func (p *queryProcessor) Run(_ context.Context) error {
	fmt.Println("Process!")

	for q := range p.recieveQueue {
		fmt.Println(q)

		token, err := parser.Tokenize(q.Task.RawQuery)
		if err != nil {
			return err
		}

		statement, err := parser.Parse(token)
		if err != nil {
			return err
		}
		_ = statement

		optimizer.Optimize()

		// TODO: 実行計画(struct?)を送るイメージ
		p.sendQueue <- q.Task.RawQuery // NOTE: 一旦、そのまま生のクエリを送る
	}
	return nil
}

func (p *queryProcessor) Close() {
	close(p.sendQueue)
}
