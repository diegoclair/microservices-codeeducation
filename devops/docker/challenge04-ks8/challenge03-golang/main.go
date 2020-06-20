package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/diegoclair/microservices-codeeducation/devops/docker/challenge04-ks8/challenge03-golang/greeting"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, greeting.greeting("Code.education Rocks!"))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
