package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Main function to start the routing server
func main() {
	router := mux.NewRouter()
	
	//API endpoint to provide min number(s)
	router.HandleFunc("/min", calculateMin).Methods(http.MethodPost)

	//API endpoint to provide max number(s)
	router.HandleFunc("/max", calculateMax).Methods(http.MethodPost)

	//API endpoint to provide average of numbers
	router.HandleFunc("/avg", calculateAvg).Methods(http.MethodPost)

	//API endpoint to provide median of numbers
	router.HandleFunc("/median", calculateMedian).Methods(http.MethodPost)

	//API endpoint to provide percentile of numbers
	router.HandleFunc("/percentile", calculatePercentile).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":3000", router))
}
