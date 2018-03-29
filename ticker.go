package main

import (
	"time"
)

const (
	interval = 2 * time.Second
	timeout  = 15 * time.Second
)

func main() {
	ticker := time.NewTicker(interval)
	timeoutCh := time.After(timeout)
	for {
		select {
		case <-ticker.C:
			// do stuff
		case <-timeoutCh:
			ticker.Stop()
			// break/return/os.exit
		}
	}

}
