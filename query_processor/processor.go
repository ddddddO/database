package query_processor

import (
	"fmt"

	"github.com/ddddddO/database/query_processor/optimizer"
	"github.com/ddddddO/database/query_processor/parser"
)

type queryProcessor struct {
	recieveQueue <-chan string
	sendQueue    chan<- string
}

func New(rQueue <-chan string, sQueue chan<- string) *queryProcessor {
	return &queryProcessor{
		recieveQueue: rQueue,
		sendQueue:    sQueue,
	}
}

func (p *queryProcessor) Run() error {
	fmt.Println("Process!")

	for q := range p.recieveQueue {
		fmt.Println(q)

		parser.Tokenize()
		parser.Parse()
		optimizer.Optimize()

		// TODO: 実行計画(struct?)を送るイメージ
		p.sendQueue <- q
	}
	return nil
}
