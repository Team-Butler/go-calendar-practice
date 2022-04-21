package main

import (
	"fmt"
	"go-calendar-practice/pkg/loaders"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	loaders.LoadApis(router)

	err := router.Run(":3000")
	if err != nil {
		fmt.Printf("Fail to start Server")
		return
	}

	fmt.Printf("Server Run")
}
