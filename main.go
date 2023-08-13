package main

import (
	"fmt"
	"miniSockS/util"
	"time"
)

func run() {
	util.DPrintf("============================\n")
	util.DPrintf("miniSockS Server is running!\n")
	util.DPrintf("============================\n")
}

func main() {
	go run()
	times := 0
	for {
		fmt.Printf("sleep times:%d\n", times)
		times++
		time.Sleep(10 * time.Second)
	}
}
