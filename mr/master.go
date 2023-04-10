package mr

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"sync"
)

type Master struct {
	sync.Mutex

	mapTasks    *Tasks
	reduceTasks *Tasks

	doneChannel chan bool
}

func (master *Master) serve() {
	rpc.Register(master)
	rpc.HandleHTTP()
	socketname := coordinatorSock()
	os.Remove(socketname)
	listner, err := net.Listen("unix", socketname)
	if err != nil {
		log.Fatal("listen error: ", err)
	}
	log.Println("Starting RPC server on: ", socketname)
	go http.Serve(listner, nil)
}

func StartMaster(files []string, nReduce int) *Master {
	mapTasks, reduceTasks := GenerateTasks(files, nReduce)
	master := new(Master)
	master.mapTasks = mapTasks
	master.reduceTasks = reduceTasks
	master.doneChannel = make(chan bool)
	master.serve()
	return master
}

func (m *Master) Wait() {
	<-m.doneChannel
}
