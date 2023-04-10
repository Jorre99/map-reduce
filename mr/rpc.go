package mr

import (
	"os"
	"strconv"
)

type RequestTaskArgs struct {
	WorkerId int
}

type RequestTaskReply struct {
	Phase           Phase
	TaskId          int
	InputFilepath   string
	TotalMapTask    int
	TotalReduceTask int
}

type ReportTaskArgs struct {
	WorkerId int
	Phase    Phase
	TaskId   int
}

type ReportTaskReply struct {
	Terminate bool
}

func coordinatorSock() string {
	s := "/var/tmp/5830-mr-"
	s += strconv.Itoa(os.Getuid())
	return s
}
