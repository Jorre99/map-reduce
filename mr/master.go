package mr

import (
	"net"
	"sync"
)

type Master struct {
	sync.Mutex

	address string

	listner  net.Listener
	shutdown chan struct{}

	doneChannel chan bool
}

func newMaster(masterAddress string) (m *Master) {
	m = new(Master)
	m.address = masterAddress
	m.doneChannel = make(chan bool)
	return
}

func StartMaster(masterAddress string) (m *Master) {
	m = newMaster(masterAddress)
	m.startRPCServer()
	return
}

func (m *Master) Wait() {
	<-m.doneChannel
}
