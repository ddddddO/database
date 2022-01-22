package main

import (
	"log"
	"time"

	"github.com/ddddddO/database/execution_engine"
	"github.com/ddddddO/database/query_processor"
	"github.com/ddddddO/database/transfer"
)

func main() {
	log.Println("start database")

	transfererToQueryprocessorQueue := make(chan string)
	transferer := transfer.New(transfererToQueryprocessorQueue)
	queryprocessorToExecutionengineQueue := make(chan string)
	queryProcessor := query_processor.New(
		transfererToQueryprocessorQueue,
		queryprocessorToExecutionengineQueue,
	)
	executionEngine := execution_engine.New(queryprocessorToExecutionengineQueue)

	go transferer.Run()
	go queryProcessor.Run()
	// NOTE: execution_engineが肝だと思う。なので、上のレイヤはモックと考えexecution_engineに重点を置くでもよさそう。
	go executionEngine.Run()

	time.Sleep(2 * time.Second)
}
