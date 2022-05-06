package main

import (
	"fmt"
)

type Alert struct {
	ID   int    `json:"id"`
	Comp string `json:"comp"`
	Time string `json:"time"`
}

func main() {
	initDBConnection()
	initHTTP()
}

func newEntry(comp, time string) {
	fmt.Println("ALERT PASSED")
	insertEntry(comp, time)
}

func check() []Alert {
	return dbcheck()
}
