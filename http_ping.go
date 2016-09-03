package main

import (
	"io"
	"net/http"
)

func httpPing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Pong")
}
