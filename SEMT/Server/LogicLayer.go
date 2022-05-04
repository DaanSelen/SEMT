package main

import (
	"fmt"
)

type Alert struct {
	ID   int    `json:"id"`
	Comp string `json:"comp"`
	Time string `json:"time"`
}

/*const (
	timeFormat = "2022-05-04 09:30:21"
)*/

func main() {
	initHTTP()
	initDBConnection()
	fmt.Println("APPLICATION IS RUNNING. Press enter to exit.")
	fmt.Scanln()
}

func newEntry(comp, time string) {
	fmt.Println("ENTRY PASSED")
	insertEntry(comp, time)
}

func check() []Alert {
	return dbcheck()
}
