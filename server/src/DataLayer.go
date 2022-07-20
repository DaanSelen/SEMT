package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var (
	entry *sql.DB
)

func initDBConnection() {
	entry, _ = sql.Open("sqlite3", "./semt.db")
	statement1, _ := entry.Prepare("CREATE TABLE IF NOT EXISTS entry (id INTEGER PRIMARY KEY, hostname varchar(100) NOT NULL, comp varchar(100) NOT NULL, time varchar(200) NOT NULL)")
	defer statement1.Close()
	log.Println("TEST")
	statement1.Exec()
	log.Println("TEST")
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

func insertEntry(hostname, comp, time string) {
	state, err := entry.Prepare("INSERT INTO entry (id, hostname, comp, time) values(null, ?, ?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer state.Close()
	_, err = state.Exec(hostname, comp, time)
	if err != nil {
		log.Println(err)
	}
	defer state.Close()
}

func dbcheck() []Alert {
	var alerts []Alert
	rows, _ := entry.Query("SELECT * FROM entry")
	defer rows.Close()
	for rows.Next() {
		var alert Alert
		rows.Scan(&alert.ID, &alert.Hostname, &alert.Comp, &alert.Time)
		alerts = append(alerts, alert)
	}
	return alerts
}
