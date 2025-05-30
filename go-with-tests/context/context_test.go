package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response string
	canceled bool
	t        *testing.T
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if s.canceled {
		s.t.Error("store was not told to cancel")
	}
}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.canceled {
		s.t.Error("store was told to cancel")
	}
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.canceled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	svr := Server(&SpyStore{response: data, canceled: false})

	req, _ := http.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	svr.ServeHTTP(resp, req)
	if resp.Body.String() != data {
		t.Errorf(`go "%s", want "%s"`, resp.Body.String(), data)
	}
}

func TestSMultipleCases(t *testing.T) {
	t.Run("returns data from store", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, canceled: false, t: t}
		svr := Server(store)

		req := httptest.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()
		svr.ServeHTTP(resp, req)

		if resp.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
		}
		store.assertWasNotCancelled()
	})
	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "hello, world"
		store := &SpyStore{response: data, t: t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		cancelCtx, cancel := context.WithCancel(context.Background())
		time.AfterFunc(100*time.Millisecond, cancel)
		request = request.WithContext(cancelCtx)

		response := httptest.NewRecorder()
		svr.ServeHTTP(response, request)

		store.assertWasCancelled()
	})
}

// V2 Google like approach

func TestSMultipleCasesV2(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {
		store := &SpyStoreV2{response: data}
		svr := ServerV2(store)
		req := httptest.NewRequest("GET", "/", nil)
		resp := httptest.NewRecorder()
		svr.ServeHTTP(resp, req)
		if resp.Body.String() != data {
			t.Errorf(`got "%s", want "%s"`, resp.Body.String(), data)
		}
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := &SpyStoreV2{response: data}
		svr := ServerV2(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if response.written {
			t.Error("a response should not have been written")
		}
	})
}
