package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/healthcheck", HealthCheck)
	fmt.Println("listening on 9997")
	http.ListenAndServe("0.0.0.0:9997", nil)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "healthy v2")
}
