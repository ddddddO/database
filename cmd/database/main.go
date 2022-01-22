package main

import (
	"log"

	"github.com/ddddddO/database/execution_engine"
	"github.com/ddddddO/database/query_processor"
	"github.com/ddddddO/database/transfer"
)

func main() {
	log.Println("start database")

	transfer.Transfer()
	query_processor.Process()
	execution_engine.Execute()
}
