package main

import (
	"fmt"
	"go-calendar-practice/pkg/loaders"
	"go-calendar-practice/pkg/middlewares/logger"

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

	server := gin.New()
	server.Use(gin.Recovery(), logger.Logger())

	loaders.LoadAPIs(server)

	return server
}
