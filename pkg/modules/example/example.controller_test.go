package example

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/magiconair/properties/assert"
)

func TestAddController(t *testing.T) {
	router := gin.New()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/example/", nil)
	router.ServeHTTP(w, req)

	// Test Add Controller
	AddController(router)

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/example/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
}
