package loaders

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestLoadAPIs(t *testing.T) {
	server := gin.Default()

	LoadAPIs(server)

	// base Group Test
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	server.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)

	// example Group Test
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/example/", nil)
	server.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}
