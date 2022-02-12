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
	sendQueue    chan<- *model.TaskAndConn
}

func New(rQueue <-chan *model.TaskAndConn, sQueue chan<- *model.TaskAndConn) *queryProcessor {
	return &queryProcessor{
		recieveQueue: rQueue,
		sendQueue:    sQueue,
	}
}

func (p *queryProcessor) Run(_ context.Context) error {
	fmt.Println("Process!")

	for q := range p.recieveQueue {
		token, err := parser.Tokenize(q.Task.RawQuery)
		if err != nil {
			return err
		}

		statement, err := parser.Parse(token)
		if err != nil {
			return err
		}

		plan, err := optimizer.Optimize(statement)
		if err != nil {
			return err
		}
		q.Task.Plan = plan

		// 実行計画をstorage engineに送る
		p.sendQueue <- q
	}
	return nil
}

func (p *queryProcessor) Close() {
	close(p.sendQueue)
}
