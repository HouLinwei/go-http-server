// test' package document.
package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestIndexHandler(t *testing.T) {

	// Create a request
	request, err := http.NewRequest("GET", "/", nil)
	//request, err := http.NewRequest("GET", "/", url.Values{
	//	"name": {"test"},
	//})
	if err != nil {
		logger.Fatal(err)
	}

	// Create a response recorder
	respRecorder := httptest.NewRecorder()

	// Create a handler
	handler := http.HandlerFunc(IndexHandler)

	// Call http handler
	handler.ServeHTTP(respRecorder, request)

	// Check status code
	if status := respRecorder.Code; status != http.StatusOK {
		logger.Fatal("bad status code: ", status)
	}

	// Checkout the response body is what we expect.
	if respRecorder.Body.String() != "hello world!" {
		logger.Fatal("bad response body: ", respRecorder.Body)
	}
}
