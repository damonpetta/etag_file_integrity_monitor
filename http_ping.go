package main

import (
	"fmt"
	"net/http"
	"os"
)

func httpPing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pong %s!", r.URL.Path[1:])
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return ":8000"
	}
	return ":" + port
}
