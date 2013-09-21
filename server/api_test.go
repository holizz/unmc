package main

import (
	"testing"
	"net/http"
	"net/http/httptest"
	"net/url"
	"github.com/stretchr/testify/assert"
	"encoding/json"
)

func TestRoot(t *testing.T) {
	record := httptest.NewRecorder()
	req := &http.Request{
		Method: "GET",
		URL: &url.URL{Path: "/"},
	}

	handleRoot(record, req)
	assert.Equal(t, record.Code, http.StatusOK)

	type Status struct {
		Status string
	}

	var status Status
	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, status.Status, "ok")
}
