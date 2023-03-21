package mr

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"os"
	"sync"
)

type Worker struct {
	sync.Mutex

	name    string
	listner net.Listener
}

func (w *Worker) register(master string) {
	debug("Register new worker \n")
	args := new(RegisterArgs)
	args.Worker = w.name
	res := call(master, "Master.Register", args, new(struct{}))
	if res == false {
		fmt.Printf("Register error %s", master)
	}
}

func StartWorker(masterAddress string, workerAddress string) {

	worker := new(Worker)
	worker.name = workerAddress
	rpcs := rpc.NewServer()
	rpcs.Register(worker)
	os.Remove(workerAddress)
	listner, err := net.Listen("unix", workerAddress)
	if err != nil {
		log.Fatal("Worker error", workerAddress, err)
	}
	worker.listner = listner
	worker.register(masterAddress)
}
