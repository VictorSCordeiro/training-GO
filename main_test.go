package main

import (
	"net/http/httptest"
	"testing"
)

func TestShowRequestPathSimple(t *testing.T) {
	// Generate a random path
	path := "randomicPathforTest999" // random string of your choice

	// Create a request with the generated path
	request := httptest.NewRequest("GET", "http://127.0.0.1:3333/"+path, nil)

	// Create a ResponseRecorder to record the response
	recordResponse := httptest.NewRecorder()

	// Call the handleRequest function with the recorder and request
	handleRequest(recordResponse, request)

	// Check if the response body is correct
	if recordResponse.Body.String() != path {
		t.Error("Test failed with path " + path)
	}
}
