package main

import (
	"fmt"
	"math"
	"net/http"
)

func sqrt() string {
	x := 0.0001

	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}

	return "Code.education Rocks"
}

func httpServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, sqrt())
}

func main() {
	http.HandleFunc("/", httpServer)
	http.ListenAndServe(":80", nil)
}
