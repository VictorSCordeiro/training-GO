package main

import (
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"
)

func TestShowRequestPathSimple(t *testing.T) {
	// Generate a random path
	path := generateRandomString(10) // random string of your choice

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

// Test case for the /prime/xxx endpoint
func TestGetNthPrime1(t *testing.T) {
	expectedPrime := 61
	req := httptest.NewRequest("GET", "http://127.0.0.1:3333/prime/18", nil)
	rec := httptest.NewRecorder()

	mainHandler := http.HandlerFunc(handlePrimeRequest)
	mainHandler.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rec.Code)
	}

	body := rec.Body.String()
	if body != strconv.Itoa(expectedPrime) {
		t.Errorf("Expected response body %d, but got %s", expectedPrime, body)
	}
}

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	randomPath := make([]byte, length)

	rand.Seed(time.Now().UnixNano())
	for i := range randomPath {
		randomPath[i] = charset[rand.Intn(len(charset))]
	}

	return string(randomPath)
}
