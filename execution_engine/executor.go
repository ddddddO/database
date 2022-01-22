package execution_engine

import (
	"fmt"

	s "github.com/ddddddO/database/execution_engine/storage_engine"
)

type executionEngine struct {
	recieveQueue <-chan string
}

func New(queue <-chan string) *executionEngine {
	return &executionEngine{
		recieveQueue: queue,
	}
}

func (e *executionEngine) Run() {
	fmt.Println("Execute!")

	for q := range e.recieveQueue {
		fmt.Println(q)
		s.TransactionManager()
		s.LockManager()
		s.AccessMethod()
		s.BufferManager()
		s.RecoveryManager()
	}
}

func (e *executionEngine) Close() {
	// noop
}
