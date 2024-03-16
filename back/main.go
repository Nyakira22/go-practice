package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type fruit struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}

var fruits = []fruit{
	{ID: 1, Name: "apple", Icon: "🍎"},
	{ID: 2, Name: "banana", Icon: "🍌"},
	{ID: 3, Name: "grape", Icon: "🍇"},
}

func main() {
	fmt.Println("Starting server at port 8080")
	http.HandleFunc("/", getFruits)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getFruits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	json.NewEncoder(w).Encode(fruits)
}