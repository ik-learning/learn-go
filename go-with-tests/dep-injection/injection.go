package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s!", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "Bonjour")
}

func main() {
	Greet(os.Stdout, "John")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(MyGreeterHandler)))
}
