# Process

A Package to check whether process is already running for Go(golang). `process` can prevent the process from being duplicated.

## Installation

``` bash
$ go get github.com/zajann/process
```

## Usage

``` go
package main

import (
	"github.com/zajann/process"
)

func main() {
    pidFilePath := "./test.pid"
    
    running, err := process.IsRunning(pidFilePath)
    if err != nil {
        panic(err)
    }
    if running {
        // Process is already running. Process exit.
    }
    
    // do something
}
```

### Return (true)

It means process is already running. You will might be shutdown the duplicated process. 

### Return (false)

It means process is not running. So you can run the process normally. If it is the first run, `process` will create PID file on the path you set.

## Example

 You can run example code in `/examples`

First, run the example code.

```bash
$ go run example1.go
```

Second, open the new shell prompt and run the example code one more. Then, you can see the message the the process is already running. Because you ran the same program over and over again.
