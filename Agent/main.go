package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type Alert struct {
	Hostname string `json:"hostname"`
	Comp     string `json:"comp"`
	Time     string `json:"time"`
}

const (
	threshold    = 20
	timeFormat   = "02-01-2006 15:04:05"
	apiServerURL = "http://192.168.10.10:2468/newentry"
)

func main() {
	fmt.Println("AGENT INITIALISING.")
	checkCpuUsage()
}

func checkCpuUsage() {
	fmt.Println("MONITORING STARTED.")
	for {
		rawPerc, _ := cpu.Percent(time.Second, false)
		cpuPerc := math.Round(rawPerc[0]*100) / 100

		fmt.Println(cpuPerc)
		if cpuPerc >= threshold {
			fmt.Println("ALERT CPU USAGE!")
			report("cpu")
		}
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
		panic(err)
	}
}
