package process

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"syscall"
)

func IsRunning(pidFilePath string) (bool, error) { // 0: not running, 1: running
	if pidFilePath == "" {
		return false, errors.New("process.checkProcessIsRunning(): invalid parmas (may be zero-values)")
	}

	if _, err := os.Stat(pidFilePath); err == nil {
		// pid file exist
		file, err := os.Open(pidFilePath)
		if err != nil {
			return false, err
		}

		var pid int
		fmt.Fscanf(file, "%d", &pid)
		file.Close()

		if process, err := os.FindProcess(pid); err == nil {
			if err := process.Signal(syscall.Signal(0)); err == nil {
				// process is running
				return true, nil
			} else {
				// process is stopped
				f, err := os.OpenFile(pidFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
				if err != nil {
					return false, err
				}
				defer f.Close()

				pid = os.Getpid()
				w := bufio.NewWriter(f)
				fmt.Fprintf(w, "%d", pid)
				w.Flush()

				return false, nil
			}
		}
	} else if os.IsNotExist(err) {
		// pid file not exist
		file, err := os.OpenFile(pidFilePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
		if err != nil {
			return false, err
		}
		defer file.Close()

		pid := os.Getpid()

		w := bufio.NewWriter(file)
		fmt.Fprintf(w, "%d", pid)
		w.Flush()

		return false, nil
	}
	return false, fmt.Errorf("process.IsRunning(): Unexpected Error")
}
