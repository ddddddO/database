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
	queryProcessor := query_processor.New(transfererToQueryprocessorQueue)

	go transferer.Run()
	go queryProcessor.Run()

	execution_engine.Execute()

	time.Sleep(2 * time.Second)
}
