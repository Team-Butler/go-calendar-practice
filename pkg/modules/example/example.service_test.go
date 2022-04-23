package example

import (
	"net/http"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGetExampleHelloWorld(t *testing.T) {
	code, message := GetExampleHelloWorld()

	assert.Equal(t, code, http.StatusOK)
	assert.Equal(t, message, "Hello, World! This Is Example Group")
}
