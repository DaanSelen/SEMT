package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dataServer = "192.168.10.15:3306"
)

var (
	entry *sql.DB
	err   error
)

func initDBConnection() {
	fmt.Println("DATABASE CONNECTION INITIALISING")
	entry, err = sql.Open("mysql", "api_server:apipassword@tcp("+dataServer+")/semt") //TO DO CHANGE AUTHENTICATION SOURCES (NOT HAVE IT HARDCODED)
	if err != nil {
		log.Fatal(err)
	} else {
		data, err := entry.Query("select count(id) from entry")
		if err != nil {
			log.Fatal(err)
		}
		defer data.Close()
		data.Next()
		var currentEntries int
		data.Scan(&currentEntries)
		fmt.Println("Current entries:", currentEntries)
	}
}

func insertEntry(hostname, comp, time string) {
	state, _ := entry.Prepare("INSERT INTO entry(hostname, comp, time) values(?, ?, ?)")
	defer state.Close()
	state.Exec(hostname, comp, time)
}

func dbcheck() []Alert {
	var alerts []Alert
	rows, _ := entry.Query("SELECT * FROM entry")
	defer rows.Close()
	for rows.Next() {
		var alert Alert
		rows.Scan(&alert.ID, &alert.Comp, &alert.Time)
		alerts = append(alerts, alert)
	}
	return alerts
}
