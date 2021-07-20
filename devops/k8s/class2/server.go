package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", Hello)
	log.Println("Server is up")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello Full Cycle K8s</h1>"))
}
