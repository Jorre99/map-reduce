package mr

import (
	"fmt"
	"net/rpc"
)

type RegisterArgs struct {
	Worker string
}

func call(srv string, rpcName string, args interface{}, reply interface{}) bool {
	conn, connErr := rpc.Dial("unix", srv)
	if connErr != nil {
		return false
	}
	defer conn.Close()

	err := conn.Call(rpcName, args, reply)
	if err == nil {
		return true
	}

	fmt.Println(err)
	return false
}
