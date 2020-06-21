package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, greeting("Code.education Rocks!"))
}

func greeting(message string) string {
	return fmt.Sprintf("<b>%s</b>", message)
}

func main() {
	http.HandleFunc("/", handler)

	log.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
