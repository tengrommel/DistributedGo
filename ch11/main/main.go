package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"DistributedGo/ch11/insertData"
)

var (
	SERVER_PORT = ":8080"
)

func insertDataEndpoint(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")
	jsonBody, _ := insertData.InsertData(r) // this method will be created in part 2
	w.Write(jsonBody)
}

/*
	end point for entering a/multple game results
	check if the request json doesn't have any errors and
	then insert the data to the database
 */

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/insert", insertDataEndpoint).Methods("POST")
	http.ListenAndServe(SERVER_PORT, router)
}
