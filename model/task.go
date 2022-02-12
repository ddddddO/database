package model

import (
	"net"
)

type TaskAndConn struct {
	Task Task
	Conn net.Conn
}

type Task struct {
	State    state  `json:"state"`
	RawQuery string `json:"raw_query"`
	Plan     Plan
}

type state uint

const (
	StateRecivedQuery state = iota
)
