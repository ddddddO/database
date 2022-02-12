package execution_engine

import (
	"context"
	"fmt"

	s "github.com/ddddddO/database/execution_engine/storage_engine"
	"github.com/ddddddO/database/model"
)

type executionEngine struct {
	recieveQueue <-chan *model.TaskAndConn
}

func New(queue <-chan *model.TaskAndConn) *executionEngine {
	return &executionEngine{
		recieveQueue: queue,
	}
}

func (e *executionEngine) Run(_ context.Context) {
	fmt.Println("Execute!")

	for q := range e.recieveQueue {
		debugRecievedQueue(q)

		s.TransactionManager()
		s.LockManager()
		s.AccessMethod()
		s.BufferManager()
		s.RecoveryManager()
	}
}

func debugRecievedQueue(t *model.TaskAndConn) {
	fmt.Println("DEBUG")
	fmt.Printf("	state: %v\n", t.Task.State)
	fmt.Printf("	query: %v\n", t.Task.RawQuery)
	fmt.Printf("	plan: %v\n", t.Task.Plan)
	fmt.Println("DEBUG")
}

func (e *executionEngine) Close() {
	// noop
}
