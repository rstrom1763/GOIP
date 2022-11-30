package main

import (
	"fmt"
	"os"
)

// This is just a placeholder for now
func main() {
	args := os.Args
	fmt.Println(args[1])
	pinger, err := ping.NewPinger("www.google.com")
	if err != nil {
		panic(err)
	}
	pinger.Count = 3
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		panic(err)
	}
	stats := pinger.Statistics() // get send/receive/duplicate/rtt stats
}
