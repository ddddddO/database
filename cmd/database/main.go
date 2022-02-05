package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddddddO/database/execution_engine"
	"github.com/ddddddO/database/model"
	"github.com/ddddddO/database/query_processor"
	"github.com/ddddddO/database/transfer"
)

// TODO: client <-> database 間のコネクションを切る
func main() {
	log.Println("start database")

	transfererToQueryprocessorQueue := make(chan *model.TaskAndConn)
	transferer := transfer.New(transfererToQueryprocessorQueue)
	queryprocessorToExecutionengineQueue := make(chan string)
	queryProcessor := query_processor.New(
		transfererToQueryprocessorQueue,
		queryprocessorToExecutionengineQueue,
	)
	executionEngine := execution_engine.New(queryprocessorToExecutionengineQueue)
	ctx, cancel := context.WithCancel(context.Background())

	go transferer.Run(ctx)
	go queryProcessor.Run(ctx)
	// NOTE: execution_engineが肝だと思う。なので、上のレイヤはモックと考えexecution_engineに重点を置くでもよさそう。
	go executionEngine.Run(ctx)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, os.Interrupt)
	<-sig

	log.Println("graceful shutdown...")
	cancel()
	transferer.Close()
	queryProcessor.Close()
	executionEngine.Close()
}
