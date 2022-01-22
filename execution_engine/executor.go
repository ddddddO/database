package execution_engine

import (
	"fmt"

	s "github.com/ddddddO/database/execution_engine/storage_engine"
)

func Execute() {
	fmt.Println("Execute!")

	s.TransactionManager()
	s.LockManager()
	s.AccessMethod()
	s.BufferManager()
	s.RecoveryManager()
}
