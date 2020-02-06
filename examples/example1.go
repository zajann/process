package main

import (
	"fmt"
	"os"
	"time"

	"github.com/zajann/process"
)

func main() {
	pidFilePath := "./test.pid"

	running, err := process.IsRunning(pidFilePath)
	if err != nil {
		panic(err)
	}

	if running {
		fmt.Println("Process is already Running...Exit")
		os.Exit(0)
	} else {
		fmt.Println("Process is not running")
	}

	for {
		time.Sleep(time.Second)
	}
}
