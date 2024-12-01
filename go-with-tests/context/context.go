package main

import (
	"context"
	"fmt"
	"net/http"
)

// https://go.dev/blog/context

func main() {
	fmt.Println("Hello World")
}

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case data := <-data:
			fmt.Fprintf(writer, data)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}

// Google Approach

// StoreV2 interface
type StoreV2 interface {
	Fetch(ctx context.Context) (string, error)
}

func ServerV2(store StoreV2) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		data, err := store.Fetch(request.Context())
		if err != nil {
			return
		}
		fmt.Fprint(writer, data)
	}
}
