package main

import (
	"fmt"
	"go-calendar-practice/pkg/loaders"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	loaders.LoadApis(r)

	r.Run(":8080")
	fmt.Printf("Server Run")
}