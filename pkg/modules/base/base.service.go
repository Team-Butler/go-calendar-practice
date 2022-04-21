package base

import "net/http"

func GetRootService() (int, string) {
	return http.StatusOK, "Hello, World! This Is Base Group"
}
