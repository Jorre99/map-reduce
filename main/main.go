package main

import (
	"fmt"
	"os"

	"github.com/jorre99/map-reduce/mr"
)

// go run main.go master localhost:7777 x1.txt .. xN.txt
// go run main.go worker localhost:7777 localhost:7778 &

func main() {
	if os.Args[1] == "master" {
		fmt.Println("Starting Master")
		var master *mr.Master
		master = mr.StartMaster("localhost:7777")
		master.Wait()

	} else {
		fmt.Println("Starting worker...")
		mr.StartWorker("localhost:7777", "localhost:7778")
		// mr.RunWorker(os.Args[2], os.Args[3], mapF, reduceF, 100, nil)
	}
}
