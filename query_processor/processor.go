package query_processor

import (
	"fmt"

	"github.com/ddddddO/database/query_processor/optimizer"
	"github.com/ddddddO/database/query_processor/parser"
)

type queryProcessor struct {
	recieveQueue <-chan string
}

func New(queue <-chan string) *queryProcessor {
	return &queryProcessor{
		recieveQueue: queue,
	}
}

func (p *queryProcessor) Run() error {
	fmt.Println("Process!")

	for q := range p.recieveQueue {
		fmt.Println(q)

		parser.Tokenize()
		parser.Parse()
		optimizer.Optimize()
	}
	return nil
}
