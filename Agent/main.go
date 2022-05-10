package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type Alert struct {
	Hostname string `json:"hostname"`
	Comp     string `json:"comp"`
	Time     string `json:"time"`
}

const (
	timeFormat = "02-01-2006 15:04:05"
)

var (
	configKeywords = []string{"THRESHOLD", "ALERTTIME", "SERVERIP"}

	count        int
	threshold    float64
	alertTime    int
	apiServerURL string
)

func main() {
	log.Println("AGENT INITIALISING.")
	initVars()
	checkCpuUsage()
}

func initVars() {
	log.Println("CHECKING CONFIG.")
	threshold, _ = strconv.ParseFloat(getInfoFromConfig(configKeywords[0]), 64)
	alertTime, _ = strconv.Atoi(getInfoFromConfig(configKeywords[1]))
	apiServerURL = "http://" + getInfoFromConfig(configKeywords[2]) + "/newentry"
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func CreateFile(name string) error {
	fo, err := os.Create(name)
	if err != nil {
		return err
	}
	defer func() {
		fo.Close()
	}()
	return nil
}

func getInfoFromConfig(keyword string) string {
	var info string
	f, err := os.Open("config.txt")
	if err != nil {
		log.Println("Opening config.txt file, perhaps there is no config.txt?\n", err)
	}
	defer f.Close()
	lineByLine := bufio.NewScanner(f)
	for lineByLine.Scan() {
		if !strings.Contains(lineByLine.Text(), "#") || lineByLine.Text() != "" { //Skipping empty rows and commented rows
			if strings.Contains(lineByLine.Text(), (keyword + " = ")) {
				info = strings.ReplaceAll(lineByLine.Text(), (keyword + " = "), "")
			}
		}
	}
	log.Println(keyword + " ASSIGNED TO: " + info)
	return info
}

func checkCpuUsage() {
	log.Println("MONITORING STARTED.")
	for {
		rawPerc, _ := cpu.Percent(time.Second, false)
		cpuPerc := math.Round(rawPerc[0]*100) / 100

		log.Println(cpuPerc)
		if cpuPerc >= threshold {
			log.Println("CPU USAGE ABOVE THRESHOLD, COUNTING.")
			countUsage("PLUS")
		} else {
			countUsage("RST")
		}
	}
}

func countUsage(command string) {
	if command == "PLUS" {
		count++
		log.Println("CURRENT COUNT:", count)
	} else if command == "RST" {
		count = 0
	}
	if count == alertTime {
		log.Println("REPORTING")
		go report("CPU")
		count = 0
	}
}

func report(comp string) {
	hostname, _ := os.Hostname()
	t := time.Now().Format(timeFormat)
	alert := Alert{
		Hostname: hostname,
		Comp:     comp,
		Time:     t,
	}
	body, _ := json.Marshal(alert)
	_, err := http.Post(apiServerURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Println("[ERROR] FAILED TO SEND HTTP REQUEST\n", err)
	}
}
