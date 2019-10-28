package main

import (
	"fmt"
	"time"

	"github.com/zajann/process"
)

func main() {
	PIDFilePath := "./test.pid"

	status, err := process.IsRunning(PIDFilePath)
	if err != nil {
		panic(err)
	}

	if status == 1 {
		fmt.Println("Process is already Running")
	} else {
		fmt.Println("Process is not running")
	}

	for {
		time.Sleep(time.Second)
	}

}
