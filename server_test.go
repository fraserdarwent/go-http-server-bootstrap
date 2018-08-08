package main

import (
	"net/http"
	"os"
	"testing"
)

const ADDRESS = "localhost"
const PORT = "8080"
const PROTOCOL = "http"

func TestMain(m *testing.M) {
	server := Server{ADDRESS + ":" + PORT}
	go server.Start()
	os.Exit(m.Run())
}

func TestHealth(t *testing.T) {
	response, err := http.Get(PROTOCOL + "://" + ADDRESS + ":" + PORT + "/health")
	if err != nil {
		t.Error(err)
	}
	if response.StatusCode != 200 {
		t.Errorf("health endpoint returned %v", response.StatusCode)
	}
}
