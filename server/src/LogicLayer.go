package main

import (
	"fmt"
)

type Alert struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	Comp     string `json:"comp"`
	Time     string `json:"time"`
}

func main() {
	go initDBConnection()
	initHTTP()
}

func newEntry(hostname, comp, time string) {
	fmt.Println("ALERT PASSED")
	insertEntry(hostname, comp, time)
}

func check() []Alert {
	return dbcheck()
}
