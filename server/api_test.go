package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func request(handler http.Handler, method string, path string) (record *httptest.ResponseRecorder) {
	record = httptest.NewRecorder()
	req := &http.Request{
		Method: method,
		URL: &url.URL{Path: path},
	}
	handler.ServeHTTP(record, req)
	return
}

func assertOK(t *testing.T, record *httptest.ResponseRecorder) {
	assert.Equal(t, record.Code, http.StatusOK)

	var status Status
	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, status.Status, "ok")
}

func TestRoot(t *testing.T) {
	mux := createMux()
	record := request(mux, "GET", "/")
	assertOK(t, record)
}

func TestList(t *testing.T) {
	mux := createMux()
	record := request(mux, "GET", "/tracks")
	assertOK(t, record)

	var status Status
	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, status.List, []Item{})
}
