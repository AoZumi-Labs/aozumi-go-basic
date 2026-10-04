package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthz(t *testing.T) {
	response := httptest.NewRecorder()
	handler("test").ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/healthz", nil))
	if response.Code != http.StatusNoContent {
		t.Fatalf("health status = %d, want %d", response.Code, http.StatusNoContent)
	}
}

func TestHomeAndUnknownRoute(t *testing.T) {
	server := httptest.NewServer(handler("AOZUMI"))
	defer server.Close()
	response, err := http.Get(server.URL + "/")
	if err != nil { t.Fatal(err) }
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil { t.Fatal(err) }
	if response.StatusCode != http.StatusOK || string(body) != "Hello AOZUMI from the AOZUMI Go fixture!\n" {
		t.Fatalf("home response = (%d, %q)", response.StatusCode, body)
	}
	response, err = http.Get(server.URL + "/missing")
	if err != nil { t.Fatal(err) }
	response.Body.Close()
	if response.StatusCode != http.StatusNotFound { t.Fatalf("unknown route status = %d, want 404", response.StatusCode) }
}
