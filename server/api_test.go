package main

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"io"
	"strings"
)

//////////////////////////////////////////////////////////////////////////////

func request(t *testing.T, handler http.Handler, method string, path string, body string) (record *httptest.ResponseRecorder) {

	var bodyReader io.Reader

	if body == "" {
		bodyReader = nil
	} else {
		bodyReader = strings.NewReader(body)
	}

	record = httptest.NewRecorder()
	req, err := http.NewRequest(method, path, bodyReader)
	assert.Nil(t, err)
	handler.ServeHTTP(record, req)
	return
}

func createHandler() (handler http.Handler) {
	initialize()
	handler = createMux()
	return
}

func assertOK(t *testing.T, record *httptest.ResponseRecorder) {
	assert.Equal(t, record.Code, http.StatusOK)

	var status Status
	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, status.Status, "ok")
}

//////////////////////////////////////////////////////////////////////////////

func TestRoot(t *testing.T) {
	h := createHandler()
	record := request(t, h, "GET", "/", "")
	assertOK(t, record)
}

func TestList(t *testing.T) {
	h := createHandler()
	record := request(t, h, "GET", "/tracks", "")
	assertOK(t, record)

	var status Status
	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)
	assert.Equal(t, status.List, []Item{})
}

func TestAdd(t *testing.T) {
	var status Status
	h := createHandler()

	// PUT a record
	record := request(t, h, "PUT", "/tracks/new", "path=/abc")
	assertOK(t, record)

	// Check that it's there
	record = request(t, h, "GET", "/tracks", "")
	assertOK(t, record)

	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)

	assert.Equal(t, len(status.List), 1)
	if len(status.List) == 1 {
		assert.Equal(t, status.List[0].Path, "/abc")
	}
}
