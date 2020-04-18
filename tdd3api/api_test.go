package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	// "net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	req := httptest.NewRequest("GET", "http://localhost:8081", nil)
	w := httptest.NewRecorder()
	health(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))
}