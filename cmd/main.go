package main

import (
	"fmt"
	"go-calendar-practice/db/ent"
	"go-calendar-practice/pkg/loaders"
	"log"

	"github.com/gin-gonic/gin"
)

const (
	DEBUG_MODE   = gin.DebugMode   // "debug"
	RELEASE_MODE = gin.ReleaseMode // "release"
)
const PORT = ":3000"

func main() {
	server := setupServer(DEBUG_MODE)
	client := setupDB()
	defer client.Close()

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

func setupDB() *ent.Client {
	client, err := loaders.PostgresSQLConnet(DEBUG_MODE)

	if err != nil {
		log.Fatalf("DB Connect Fail %v", err)
	}
	
	if err != nil {
		log.Fatalf("fail %v", err)
	}
	
	// DB 사전 세팅 코드
	// ctx := context.Background()
	// client.Schema.Create(ctx, migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	// client.User.Create().SetID(1).SetName("test").SaveX(context.TODO())

	return client
}