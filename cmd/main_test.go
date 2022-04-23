package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSetupServer(t *testing.T) {
	// Server Setup
	server := setupServer()

	// response test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
