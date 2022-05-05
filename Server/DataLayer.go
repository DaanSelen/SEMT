package main

import (
	"database/sql"
	"fmt"

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
		panic("DATABASE CONNECTION FAILED")
	} else {
		data, _ := entry.Query("select count(id) from entry")
		defer data.Close()
		data.Next()
		var currentEntries int
		data.Scan(&currentEntries)
		fmt.Println("Current entries:", currentEntries)
	}
}

func insertEntry(comp, time string) {
	state, _ := entry.Prepare("INSERT INTO entry(comp, time) values(?, ?)")
	defer state.Close()
	state.Exec(comp, time)
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
