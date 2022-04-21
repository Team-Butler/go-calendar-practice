package example

import "net/http"

func GetExampleHelloWorld() (int, string) {
	return http.StatusOK, "Hello, World!"
}
