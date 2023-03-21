package mr

import (
	"log"
	"net"
	"net/rpc"
	"os"
)

func (master *Master) startRPCServer() {
	rpcs := rpc.NewServer()
	rpcs.Register(master)
	os.Remove(master.address)
	listner, err := net.Listen("unix", master.address)
	if err != nil {
		log.Fatal("Server", master.address, " err: ", err)
	}
	master.listner = listner

	// now that we are listening on the master address, can fork off
	// accepting connections to another thread.
	go func() {
	loop:
		for {
			select {
			case <-master.shutdown:
				break loop
			default:
			}
			conn, err := master.listner.Accept()
			if err == nil {
				go func() {
					rpcs.ServeConn(conn)
					conn.Close()
				}()
			} else {
				debug("RegistrationServer: accept error\n")
				break
			}
		}
		debug("RegistrationServer: done\n")
	}()
}
