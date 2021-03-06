package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func initHTTP() {
	fmt.Println("API SERVER INITIALISING")
	semt := mux.NewRouter().StrictSlash(true)

	semt.HandleFunc("/", handleRootEndpoint).Methods("GET")
	semt.HandleFunc("/monitor/uptime", handleMonitorUptimeEndpoint).Methods("PATCH")
	semt.HandleFunc("/monitor/cpu", handleMonitorCPUAlertEndpoint).Methods("POST")
	semt.HandleFunc("/monitor/ram", handleMonitorRAMAlertEndpoint).Methods("POST")
	semt.HandleFunc("/control/uptime", handleControlUptimeEndpoint).Methods("GET")
	semt.HandleFunc("/control/cpu", handleControlCPUEndpoint).Methods("GET")
	semt.HandleFunc("/control/ram", handleControlRAMEndpoint).Methods("GET")

	http.ListenAndServe((":2468"), semt)
}

func handleRootEndpoint(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	json.NewEncoder(w).Encode("Root directory endpoint hit! No options here. Check /newentry or /check")
}

func handleMonitorUptimeEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func handleMonitorRAMAlertEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

}

func handleMonitorCPUAlertEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var alert Alert
	_ = json.NewDecoder(r.Body).Decode(&alert)
	newEntry(alert.Hostname, alert.Comp, alert.Time)
}

func handleControlUptimeEndpoint(w http.ResponseWriter, r *http.Request) {

}

func handleControlCPUEndpoint(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	alters := check()
	if alters == nil {
		json.NewEncoder(w).Encode("THERE ARE NO ENTRIES IN THE DATABASE")
	} else {
		json.NewEncoder(w).Encode(alters)
	}
}

func handleControlRAMEndpoint(w http.ResponseWriter, r *http.Request) {

}
