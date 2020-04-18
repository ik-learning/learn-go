package main

import (
	"io"
	"net/http"
)

// ItemList list of item
type ItemList []struct {
	Item string `json:"item"`
}

func main() {
	http.HandleFunc("/health", health)
	http.ListenAndServe(":8080", nil)
}

// health provide base health provide
func health(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "{\"message\":\"OK\"}\n")
}