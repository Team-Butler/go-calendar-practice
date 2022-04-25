package base

import (
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetRootService(t *testing.T) {
	code, message := GetRootService()

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, message, "Hello, World! This Is Base Group")
}
