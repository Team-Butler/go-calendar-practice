package main

import (
	"fmt"
	"go-calendar-practice/pkg/loaders"

	"github.com/gin-gonic/gin"
)

func main() {
	server := setupServer()

	err := server.Run(":3000")
	if err != nil {
		fmt.Printf("Fail to start Server")
		return
	}

	fmt.Printf("Server Run")
}

func setupServer() *gin.Engine {
	router := gin.Default()

	loaders.LoadAPIs(router)

	return router
}
