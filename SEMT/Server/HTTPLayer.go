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

	go http.ListenAndServe((":2468"), semt)
}

func handleRootEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	json.NewEncoder(w).Encode("Root directory endpoint hit! No options here.")
}

func handleNewEntryEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	compQuery, ok1 := r.URL.Query()["comp"]
	timeQuery, ok2 := r.URL.Query()["time"]

	switch {
	case (ok1 && len(compQuery) > 0) && (ok2 && len(timeQuery) > 0): //ALL PARAMETERS ARE MET
		/*if correct := */ newEntry(compQuery[0], timeQuery[0]) /*; correct { 																//TO DO FILTER INPUT

		}*/
	case (!ok1 || len(compQuery) <= 0) && (ok2 && len(timeQuery) > 0): //ONLY TIME PARAMETER IS MET TYPE IS MISSING
		fallthrough
	case (!ok2 || len(timeQuery) <= 0) && (ok1 && len(compQuery) > 0): //ONLY TYPE PARAMETER IS MET TIME IS MISSING
		fallthrough
	case (!ok1 || len(timeQuery) <= 0) && (!ok1 || len(compQuery) <= 0): //BOTH PARAMETERS ARE MISSING
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Either 1 or both parameters are missing: type and/or time")
	}
}

func handleCheckEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	alters := check()
	json.NewEncoder(w).Encode(alters)
}
