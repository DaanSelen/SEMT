package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Alert struct {
	ID       int    `json:"id"`
	Hostname string `json:"hostname"`
	Comp     string `json:"comp"`
	Time     string `json:"time"`
}

var (
	configKeywords = []string{"DATASERVERIP", "DATABASEUSERNAME", "DATABASEPASSWORD"}

	dataServer string
	dataUser   string
	dataPass   string
)

func main() {
	initVars()
	go initDBConnection()
	initHTTP()
}

func initVars() {
	dataServer = getInfoFromConfig(configKeywords[0])
	dataUser = getInfoFromConfig(configKeywords[1])
	dataPass = getInfoFromConfig(configKeywords[2])
}

func getInfoFromConfig(keyword string) string {
	var info string
	f, err := os.Open("config.txt")
	if err != nil {
		log.Println("Opening config.txt file, perhaps there is no config.txt?\n", err)
	}
	defer f.Close()
	for lineByLine := bufio.NewScanner(f); lineByLine.Scan(); {
		if !(strings.Contains(lineByLine.Text(), "#") || lineByLine.Text() == "") && strings.Contains(lineByLine.Text(), (keyword+" = ")) { //Skipping empty rows and commented rows and checking for the ' = ' in the config.
			info = strings.ReplaceAll(lineByLine.Text(), (keyword + " = "), "") //Getting the important information from the line with the keyword.
		}
	}
	log.Println(keyword + " ASSIGNED TO: " + info)
	return info
}

func newEntry(hostname, comp, time string) {
	fmt.Println("ALERT PASSED")
	insertEntry(hostname, comp, time)
}

func check() []Alert {
	return dbcheck()
}
