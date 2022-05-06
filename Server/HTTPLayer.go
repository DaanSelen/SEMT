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
	semt.HandleFunc("/newentry", handleNewEntryEndpoint).Methods("POST")
	semt.HandleFunc("/check", handleCheckEndpoint).Methods("GET")

	http.ListenAndServe((":2468"), semt)
}

func handleRootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	json.NewEncoder(w).Encode("Root directory endpoint hit! No options here.")
}

func handleNewEntryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var alert Alert
	_ = json.NewDecoder(r.Body).Decode(&alert)
	newEntry(alert.Comp, alert.Time)
}

func handleCheckEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alters := check()
	json.NewEncoder(w).Encode(alters)
}
