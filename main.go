package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/*
	Creating country struct
*/
type Country struct {
	CName   string `json:"cName"`
	Capital string `json:"capital"`
}

type Countries []Country

// Show home page
func homePage(w http.ResponseWriter, r *http.Request){
	_, _ = fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Request coming from client will call the required function
func handleRequests() {
	http.HandleFunc("/home", homePage)
	http.HandleFunc("/countries", returnAllCountries)
	err := http.ListenAndServe(":8080", nil)
	log.Fatal(err)
}

// List out the country added in list
func returnAllCountries(w http.ResponseWriter, r *http.Request){
	countries := Countries{
		Country{CName: "India", Capital: "Delhi"},
		Country{CName: "Nepal", Capital: "Kathmandu"},
	}
	fmt.Println("Endpoint Hit: returnAllCountries")

	_ = json.NewEncoder(w).Encode(countries)
}

func main() {
	handleRequests()
}