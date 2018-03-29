package main

import (
	"time"
)

const (
	interval = 5 * time.Second
	timeout  = 5 * time.Minute
)

func main() {
	ticker := time.NewTicker(interval)
	for {
		select {
		case <-ticker.C:
			// do stuff
		case <-time.After(timeout):
			ticker.Stop()
			// break/return/os.exit
		}
	}

}
