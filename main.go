package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/corona/all-country", getAllCountryData).Methods("GET")
	router.HandleFunc("/corona/by-country/{countryCode}", getByCountry).Methods("GET")
	router.HandleFunc("/corona/global", getGlobalData).Methods("GET")
	router.HandleFunc("/corona/highest-affected-country", getByHighest).Methods("GET")
	router.HandleFunc("/corona/estimation-on-x-day/{countryCode}/{chance}/{person}/{day}", getAverageOnXDay).Methods("GET")

	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))
}
