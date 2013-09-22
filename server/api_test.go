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

	if body != "" {
		req.Header.Set("Content-Type", "application/x-www-form-encoded")
	}

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
	//TODO: revert back to having the args in the body of the request
	// record := request(t, h, "PUT", "/tracks/new", "path=/abc")
	record := request(t, h, "PUT", "/tracks/new?path=/abc", "")
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

func TestAddTwo(t *testing.T) {
	var status Status
	h := createHandler()

	// PUT some records
	record := request(t, h, "PUT", "/tracks/new?path=/abc", "")
	assertOK(t, record)
	err := json.Unmarshal(record.Body.Bytes(), &status)
	id1 := status.Id

	record = request(t, h, "PUT", "/tracks/new?path=/def", "")
	assertOK(t, record)
	err = json.Unmarshal(record.Body.Bytes(), &status)
	id2 := status.Id

	// Check the ids are different
	assert.NotEqual(t, id1, id2)

	// Check that they're there
	record = request(t, h, "GET", "/tracks", "")
	assertOK(t, record)

	err = json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)

	assert.Equal(t, len(status.List), 2)
	if len(status.List) == 2 {
		assert.Equal(t, status.List[0].Id, id1)
		assert.Equal(t, status.List[0].Path, "/abc")
		assert.Equal(t, status.List[1].Id, id2)
		assert.Equal(t, status.List[1].Path, "/def")
	}
}

func TestDelete(t *testing.T) {
	var status Status
	h := createHandler()

	// It has some items already
	items = []Item{
		Item{
			Id: 1,
			Path: "/123",
		},
		Item{
			Id: 2,
			Path: "/456",
		},
		Item{
			Id: 3,
			Path: "/789",
		},
	}

	// DELETE a record
	record := request(t, h, "DELETE", "/tracks/2", "")
	assertOK(t, record)

	// Check that two items are still there
	record = request(t, h, "GET", "/tracks", "")
	assertOK(t, record)

	err := json.Unmarshal(record.Body.Bytes(), &status)
	assert.Nil(t, err)

	assert.Equal(t, len(status.List), 2)
	if len(status.List) == 2 {
		assert.Equal(t, status.List[0].Id, 1)
		assert.Equal(t, status.List[0].Path, "/123")
		assert.Equal(t, status.List[1].Id, 3)
		assert.Equal(t, status.List[1].Path, "/789")
	}
}
