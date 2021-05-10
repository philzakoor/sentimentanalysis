package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//function for the user help in the API from the "/" endpoint
func index (w http.ResponseWriter, r *http.Request) {
	_,err := fmt.Fprintln(w, "How to use")
	_,err = fmt.Fprintln(w, "type into browser localhost:8080/sentiment&filter=")
	_,err = fmt.Fprintln(w, "Then the topic you wish to study ")
	if err != nil {
		log.Println("error on write")
	}
}

//function for the sentiment in the API from the "/sentiment/{filter}" endpoint
func sentiment (w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		//w.Header().Set("Access-Control-Allow-Headers","Content-Type,access-control-allow-origin, access-control-allow-headers")

		filter := mux.Vars(r)["filter"]

	a := analyze(filter)

	err := json.NewEncoder(w).Encode(a)

	if err != nil {
		log.Println(err)
	}
}

//starts the router and server
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", index)
	router.HandleFunc("/sentiment/{filter}", sentiment)
	log.Fatal(http.ListenAndServe(":8080", router))
}
