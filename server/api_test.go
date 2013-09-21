package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func request(handler func(http.ResponseWriter, *http.Request), method string, path string) (record *httptest.ResponseRecorder) {
	record = httptest.NewRecorder()
	req := &http.Request{
		Method: method,
		URL: &url.URL{Path: path},
	}
	handler(record, req)
	return
}

func assertOK(t *testing.T, record *httptest.ResponseRecorder) {
	assert.Equal(t, record.Code, http.StatusOK)

	type Status struct {
		Status string
	}

	var status Status
	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, status.Status, "ok")
}

func TestRoot(t *testing.T) {
	record := request(handleRoot, "GET", "/")
	assertOK(t, record)
}
