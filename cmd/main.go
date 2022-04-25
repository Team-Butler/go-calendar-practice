package main

import (
	"fmt"
	"go-calendar-practice/pkg/loaders"

	"github.com/gin-gonic/gin"
)

const (
	DEBUG_MODE   = gin.DebugMode   // "debug"
	RELEASE_MODE = gin.ReleaseMode // "release"
)
const PORT = ":3000"

func main() {

	server := setupServer(DEBUG_MODE)

	err := server.Run(PORT)
	if err != nil {
		fmt.Printf("Fail to start Server")
		return
	}

	fmt.Printf("Server Run")
}

func setupServer(mode string) *gin.Engine {
	gin.SetMode(mode)

	router := gin.Default()

	loaders.LoadAPIs(router)

	return router
}
